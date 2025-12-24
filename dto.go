package gomcsmp

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/netip"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type PlayerRegistry struct {
	data map[string]*uuid.UUID
}

func newPlayerRegistry(len int) *PlayerRegistry {
	return &PlayerRegistry{
		data: make(map[string]*uuid.UUID, len),
	}
}

func NewPlayerRegistry(players []Player) *PlayerRegistry {
	r := &PlayerRegistry{
		data: make(map[string]*uuid.UUID, len(players)),
	}

	for _, p := range players {
		if p.Name == "" {
			continue
		}
		r.data[p.Name] = p.ID
	}

	return r
}

func NewPlayerRegistryNames(name ...string) *PlayerRegistry {
	r := newPlayerRegistry(len(name))

	for _, p := range name {
		if p == "" {
			continue
		}
		r.data[p] = nil
	}

	return r
}

func (r *PlayerRegistry) Players() []Player {
	if r == nil || len(r.data) == 0 {
		return nil
	}

	players := make([]Player, 0, len(r.data))

	for name, id := range r.data {
		players = append(players, Player{
			Name: name,
			ID:   id,
		})
	}

	return players
}

func (r *PlayerRegistry) IDs() []uuid.UUID {
	if r == nil || len(r.data) == 0 {
		return nil
	}

	ids := []uuid.UUID{}
	for _, id := range r.data {
		if id != nil && *id != uuid.Nil {
			ids = append(ids, *id)
		}
	}
	return ids
}

func (r *PlayerRegistry) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Players())
}

func (r *PlayerRegistry) UnmarshalJSON(b []byte) error {
	var players []Player
	if err := json.Unmarshal(b, &players); err != nil {
		return err
	}

	r.data = make(map[string]*uuid.UUID, len(players))
	for _, p := range players {
		r.data[p.Name] = p.ID
	}

	return nil
}

func (r *PlayerRegistry) Filter(f func(data Player) bool) *PlayerRegistry {
	newR := newPlayerRegistry(0)
	for _, player := range r.Players() {
		if f(player) {
			newR.Add(player)
		}
	}
	return newR
}

// ============

func (r *PlayerRegistry) Add(player Player) {
	r.data[player.Name] = player.ID
}

func (r *PlayerRegistry) Online() int {
	return len(r.data)
}

func (r *PlayerRegistry) Contains(name string) bool {
	_, ok := r.data[name]
	return ok
}

func (r *PlayerRegistry) UUIDByName(name string) (id uuid.UUID, ok bool) {
	srcId, ok := r.data[name]
	if !ok {
		return uuid.Nil, false
	}

	if srcId == nil {
		return uuid.Nil, false
	}

	return *srcId, true
}

func (r *PlayerRegistry) PlayerByName(name string) (player Player, ok bool) {
	srcId, ok := r.data[name]
	if !ok {
		return Player{}, false
	}

	player = Player{
		Name: name,
		ID:   srcId,
	}

	return player, true
}

// ============

// Player - represents a Minecraft player.
// Name is the player's username, ID is the optional UUID of the player.
type Player struct {
	Name string     `json:"name"`
	ID   *uuid.UUID `json:"id,omitempty"`
}

// NewPlayer - creates a new Player instance with the given name.
// The UUID is left nil and can be set later.
func NewPlayer(name string) Player {
	return Player{
		Name: name,
	}
}

func (player Player) Kick(literal string, translatable string, params ...string) KickPlayer {
	return NewKickPlayer(player.Name, NewMessage(literal, translatable, params...))
}

func (player Player) SystemMessageOverlay(literal string, translatable string, params ...string) SystemMessage {
	p := newPlayerRegistry(1)
	p.Add(player)
	return SystemMessage{
		ReceivingPlayers: p,
		Overlay:          true,
		Message:          NewMessage(literal, translatable, params...),
	}
}

func (player Player) SystemMessage(literal string, translatable string, params ...string) SystemMessage {
	p := newPlayerRegistry(1)
	p.Add(player)
	return SystemMessage{
		ReceivingPlayers: p,
		Overlay:          false,
		Message:          NewMessage(literal, translatable, params...),
	}
}

// ============

type Operator struct {
	Permission          int    `json:"permissionLevel"`
	BypassesPlayerLimit bool   `json:"bypassesPlayerLimit"`
	Player              Player `json:"player"`
}

func NewOperator(player Player, perms int, bypassLimit bool) Operator {
	return Operator{
		Permission:          perms,
		BypassesPlayerLimit: bypassLimit,
		Player:              player,
	}
}

// ============

type GameRuleType string

const (
	UntypedGameRule GameRuleType = ""
	IntegerGameRule GameRuleType = "integer"
	BooleanGameRule GameRuleType = "boolean"
)

// GameRule - represents a game rule key-value pair in the server.
type GameRule struct {
	Type  GameRuleType `json:"type,omitempty"`
	Value string       `json:"value"`
	Key   string       `json:"key"`
}

// NewGameRule - creates a new GameRule with the given key and value.
func NewGameRule(value, key string, ruleType GameRuleType) GameRule {
	return GameRule{
		Type:  ruleType,
		Value: value,
		Key:   key,
	}
}

func NewGameRuleBoolean(value bool, key string) GameRule {
	return NewGameRule(strconv.FormatBool(value), key, BooleanGameRule)
}

func NewGameRuleInteger(value int, key string) GameRule {
	return NewGameRule(strconv.Itoa(value), key, IntegerGameRule)
}

func (gr GameRule) Untyped() bool {
	return gr.Type == UntypedGameRule
}

func (gr GameRule) RuleKey() string {
	return gr.Key
}

func (gr GameRule) RuleValue() any {
	return gr.Value
}

func (gr GameRule) Boolean() (bool, error) {
	if gr.Type != BooleanGameRule {
		return false, errors.New("not boolean rule")
	}
	return gr.Value == "true", nil
}

func (gr GameRule) Integer() (int, error) {
	if gr.Type != IntegerGameRule {
		return 0, errors.New("not integer rule")
	}

	v, err := strconv.Atoi(gr.Value)
	if err != nil {
		return 0, fmt.Errorf("invalid value: %w", err)
	}

	return v, nil
}

// ============

// IncomingIPBan - represents an incoming IP ban request or record.
// Contains the reason, expiration time, IP address, source of ban, and the associated player.
type IncomingIPBan struct {
	Reason  string `json:"reason"`
	Expires string `json:"expires"`
	IP      string `json:"ip"`
	Source  string `json:"source"`
	Player  Player `json:"player"`
}

// Addr - parses the IP field into a netip.Addr.
// Returns an error if the IP is invalid.
func (iipb IncomingIPBan) Addr() (netip.Addr, error) {
	a, err := netip.ParseAddr(iipb.IP)
	if err != nil {
		return netip.Addr{}, fmt.Errorf("invalid ban IP: %w", err)
	}
	return a, nil
}

// NetIP - parses the IP field into a net.IP.
// Returns an error if the IP is invalid.
func (iipb IncomingIPBan) NetIP() (net.IP, error) {
	a := net.ParseIP(iipb.IP)
	if a == nil {
		return nil, errors.New("invalid ban IP")
	}
	return a, nil
}

// ExpireIn - parses the Expires field into a time.Time object.
// Returns zero time (time.Time{}) if parsing fails.
func (iipb IncomingIPBan) ExpireIn() time.Time {
	t, err := time.Parse(time.ANSIC, iipb.Expires)
	if err != nil {
		return time.Time{}
	}
	return t
}

// Expired - checks if the ban has already expired.
// Returns true if Expires is valid and the current time is after it.
func (iipb IncomingIPBan) Expired() bool {
	t := iipb.ExpireIn()
	return !t.IsZero() && time.Now().After(t)
}

// ============

type Message struct {
	Translatable       string   `json:"translatable"`
	TranslatableParams []string `json:"translatableParams"`
	Literal            string   `json:"literal"`
}

func NewMessage(literal string, translatable string, params ...string) Message {
	return Message{
		Literal:            literal,
		Translatable:       translatable,
		TranslatableParams: params,
	}
}

// ============

type SystemMessage struct {
	ReceivingPlayers *PlayerRegistry `json:"receivingPlayers"`
	Overlay          bool            `json:"overlay"`
	Message          Message         `json:"message"`
}

// ============

type KickPlayer struct {
	Player  Player  `json:"player"`
	Message Message `json:"message"`
}

func NewKickPlayer(name string, msg Message) KickPlayer {
	return KickPlayer{
		Player:  NewPlayer(name),
		Message: msg,
	}
}

// ============

// IPBan - represents an IP ban record.
// Contains the reason, expiration time, IP address, and the source of the ban.
type IPBan struct {
	Reason  string `json:"reason"`
	Expires string `json:"expires"`
	IP      string `json:"ip"`
	Source  string `json:"source"`
}

// NewIPBan - creates a new IPBan with the given parameters.
func NewIPBan(ip net.IP, expires time.Time, reason, source string) (IPBan, error) {

	if expires.Before(time.Now()) {
		return IPBan{}, errors.New("invalid expiration time")
	}

	return IPBan{
		IP:      ip.String(),
		Reason:  reason,
		Expires: expires.Format(time.ANSIC),
		Source:  source,
	}, nil
}

// Addr - parses the IP field into a netip.Addr.
// Returns an error if the IP is invalid.
func (ipb IPBan) Addr() (netip.Addr, error) {
	a, err := netip.ParseAddr(ipb.IP)
	if err != nil {
		return netip.Addr{}, fmt.Errorf("invalid ban IP: %w", err)
	}
	return a, nil
}

// NetIP - parses the IP field into a net.IP.
// Returns an error if the IP is invalid.
func (ipb IPBan) NetIP() (net.IP, error) {
	a := net.ParseIP(ipb.IP)
	if a == nil {
		return nil, errors.New("invalid ban IP")
	}
	return a, nil
}

// ExpireIn - parses the Expires field into a time.Time object.
// Returns zero time (time.Time{}) if parsing fails.
func (ipb IPBan) ExpireIn() time.Time {
	t, err := time.Parse(time.ANSIC, ipb.Expires)
	if err != nil {
		return time.Time{}
	}
	return t
}

// Expired - checks if the ban has already expired.
// Returns true if Expires is valid and the current time is after it.
func (ipb IPBan) Expired() bool {
	t := ipb.ExpireIn()
	return !t.IsZero() && time.Now().After(t)
}

// ============

// UserBan - represents a user ban record.
// Contains the reason, expiration time, source of the ban, and the banned player.
type UserBan struct {
	Reason  string `json:"reason"`
	Expires string `json:"expires"`
	Source  string `json:"source"`
	Player  Player `json:"player"`
}

// NewUserBan - creates a new UserBan with the given parameters.
func NewUserBan(player Player, expires time.Time, reason, source string) (UserBan, error) {

	if expires.Before(time.Now()) {
		return UserBan{}, errors.New("invalid expiration time")
	}

	return UserBan{
		Player:  player,
		Reason:  reason,
		Expires: expires.Format(time.ANSIC),
		Source:  source,
	}, nil
}

// ExpireIn - parses the Expires field into a time.Time object.
// Returns zero time (time.Time{}) if parsing fails.
func (ub UserBan) ExpireIn() time.Time {
	t, err := time.Parse(time.ANSIC, ub.Expires)
	if err != nil {
		return time.Time{}
	}
	return t
}

// Expired - checks if the user ban has already expired.
// Returns true if Expires is valid and the current time is after it.
func (ub UserBan) Expired() bool {
	t := ub.ExpireIn()
	return !t.IsZero() && time.Now().After(t)
}

// ============

type Version struct {
	Protocol int    `json:"protocol"`
	Name     string `json:"name"`
}

func parseSub(s string) (v int, ok bool) {
	av, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}
	return av, true
}

func (v Version) VersionNumbers() (major, minor, patch int) {
	sub := strings.Split(v.Name, ".")
	if len(sub) != 3 {
		return 0, 0, 0
	}

	var ok bool

	if major, ok = parseSub(sub[0]); !ok {
		return 0, 0, 0
	}
	if minor, ok = parseSub(sub[1]); !ok {
		return 0, 0, 0
	}
	if patch, ok = parseSub(sub[2]); !ok {
		return 0, 0, 0
	}

	return major, minor, patch
}

// ============

type ServerState struct {
	Players *PlayerRegistry `json:"players"`
	Started bool            `json:"started"`
	Version Version         `json:"version"`
}
