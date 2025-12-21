package gomcsmp

import (
	"errors"
	"fmt"
	"net"
	"net/netip"
	"strconv"
	"time"

	"github.com/google/uuid"
)

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
func NewGameRule(value, key string) GameRule {
	return GameRule{
		Type:  UntypedGameRule,
		Value: value,
		Key:   key,
	}
}

func NewGameRuleTyped(value, key string, ruleType GameRuleType) GameRule {
	if ruleType == UntypedGameRule {
		return NewGameRule(value, key)
	}

	return GameRule{
		Type:  ruleType,
		Value: value,
		Key:   key,
	}
}

func (gr GameRule) Untyped() bool {
	return gr.Type == UntypedGameRule
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
	ReceivingPlayers []Player `json:"receivingPlayers"`
	Overlay          bool     `json:"overlay"`
	Message          Message  `json:"message"`
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

// ============

type ServerState struct {
	Players []Player `json:"players"`
	Started bool     `json:"started"`
	Version Version  `json:"version"`
}

func (ss ServerState) Online() int {
	return len(ss.Players)
}

func (ss ServerState) IsOnline(player Player) bool {
	for _, onlinePlayer := range ss.Players {
		if onlinePlayer.Name == player.Name {
			return true
		}
	}
	return false
}
