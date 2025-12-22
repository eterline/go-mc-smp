package gomcsmp

import (
	"context"
	"time"

	"github.com/eterline/go-mc-smp/internal/jsonrpc"
	"github.com/eterline/go-mc-smp/internal/usage"
)

// Endpoints are accessible at minecraft:serversettings
//
// Path                              | Description
// ----------------------------------|-------------------------------------------------------------
// /autosave                         | Get whether automatic world saving is enabled
// /autosave/set                     | Enable or disable automatic world saving
// /difficulty                       | Get the current server difficulty
// /difficulty/set                   | Set the server difficulty
// /enforce_allowlist                | Get whether allowlist enforcement is enabled
// /enforce_allowlist/set            | Enable or disable allowlist enforcement
// /use_allowlist                    | Get whether the allowlist is enabled
// /use_allowlist/set                | Enable or disable the allowlist
// /max_players                      | Get the maximum number of players
// /max_players/set                  | Set the maximum number of players
// /pause_when_empty_seconds         | Get auto-pause delay when server is empty (seconds)
// /pause_when_empty_seconds/set     | Set auto-pause delay when server is empty
// /player_idle_timeout              | Get idle player kick timeout (seconds)
// /player_idle_timeout/set          | Set idle player kick timeout
// /allow_flight                     | Get whether flight is allowed in Survival mode
// /allow_flight/set                 | Enable or disable flight in Survival mode
// /motd                             | Get the server MOTD
// /motd/set                         | Set the server MOTD
// /spawn_protection_radius          | Get spawn protection radius (blocks)
// /spawn_protection_radius/set      | Set spawn protection radius
// /force_game_mode                  | Get whether default game mode is forced
// /force_game_mode/set              | Enable or disable forced game mode
// /game_mode                        | Get the default game mode
// /game_mode/set                    | Set the default game mode
// /view_distance                    | Get view distance (chunks)
// /view_distance/set                | Set view distance
// /simulation_distance              | Get simulation distance (chunks)
// /simulation_distance/set          | Set simulation distance
// /accept_transfers                 | Get whether player transfers are accepted
// /accept_transfers/set             | Enable or disable player transfers
// /status_heartbeat_interval        | Get status heartbeat interval (seconds)
// /status_heartbeat_interval/set    | Set status heartbeat interval
// /operator_user_permission_level   | Get operator permission level
// /operator_user_permission_level/set | Set operator permission level
// /hide_online_players              | Get whether online players are hidden from status
// /hide_online_players/set          | Enable or disable hiding online players
// /status_replies                   | Get whether status replies are enabled
// /status_replies/set               | Enable or disable status replies
// /entity_broadcast_range           | Get entity broadcast range (percentage)
// /entity_broadcast_range/set       | Set entity broadcast range (percentage)

// ===========

// SettingsAutosave - Get whether automatic world saving is enabled on the server
func (rpc *RPCClient) SettingsAutosave(ctx context.Context) (bool, error) {
	method := usage.NewMethod("serversettings").Add("autosave").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return false, err
	}

	data, err := jsonrpc.DecodeRPCResult[bool](r)
	if err != nil {
		return false, err
	}

	return *data, nil
}

// SettingsAutosaveSet - Enable or disable automatic world saving on the server
func (rpc *RPCClient) SettingsAutosaveSet(ctx context.Context, enable bool) (bool, error) {
	method := usage.NewMethod("serversettings").Add("autosave").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, enable)
	if err != nil {
		return false, err
	}

	data, err := jsonrpc.DecodeRPCResult[bool](r)
	if err != nil {
		return false, err
	}

	return *data, nil
}

// ===========

// SettingsDifficulty - Get the current difficulty level of the server
func (rpc *RPCClient) SettingsDifficulty(ctx context.Context) (string, error) {
	method := usage.NewMethod("serversettings").Add("difficulty").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return "", err
	}

	data, err := jsonrpc.DecodeRPCResult[string](r)
	if err != nil {
		return "", err
	}

	return *data, nil
}

// SettingsDifficultySet - Set the difficulty level of the server
func (rpc *RPCClient) SettingsDifficultySet(ctx context.Context, difficulty string) (string, error) {
	method := usage.NewMethod("serversettings").Add("difficulty").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, difficulty)
	if err != nil {
		return "", err
	}

	data, err := jsonrpc.DecodeRPCResult[string](r)
	if err != nil {
		return "", err
	}

	return *data, nil
}

// ===========

// SettingsEnforceAllowlist - Get whether allowlist enforcement is enabled (kicks players immediately when removed from allowlist)
func (rpc *RPCClient) SettingsEnforceAllowlist(ctx context.Context) (bool, error) {
	method := usage.NewMethod("serversettings").Add("enforce_allowlist").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return false, err
	}

	data, err := jsonrpc.DecodeRPCResult[bool](r)
	if err != nil {
		return false, err
	}

	return *data, nil
}

// SettingsEnforceAllowlistSet - Enable or disable allowlist enforcement (when enabled, players are kicked immediately upon removal from allowlist)
func (rpc *RPCClient) SettingsEnforceAllowlistSet(ctx context.Context, enforce bool) (bool, error) {
	method := usage.NewMethod("serversettings").Add("enforce_allowlist").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, enforce)
	if err != nil {
		return false, err
	}

	data, err := jsonrpc.DecodeRPCResult[bool](r)
	if err != nil {
		return false, err
	}

	return *data, nil
}

// ===========

// SettingsUseAllowlist - Get whether the allowlist is enabled on the server
func (rpc *RPCClient) SettingsUseAllowlist(ctx context.Context) (bool, error) {
	method := usage.NewMethod("serversettings").Add("use_allowlist").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return false, err
	}

	data, err := jsonrpc.DecodeRPCResult[bool](r)
	if err != nil {
		return false, err
	}

	return *data, nil
}

// SettingsUseAllowlistSet - Enable or disable the allowlist on the server (controls whether only allowlisted players can join)
func (rpc *RPCClient) SettingsUseAllowlistSet(ctx context.Context, use bool) (bool, error) {
	method := usage.NewMethod("serversettings").Add("use_allowlist").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, use)
	if err != nil {
		return false, err
	}

	data, err := jsonrpc.DecodeRPCResult[bool](r)
	if err != nil {
		return false, err
	}

	return *data, nil
}

// ===========

// SettingsMaxPlayers - Get the maximum number of players allowed to connect to the server
func (rpc *RPCClient) SettingsMaxPlayers(ctx context.Context) (int, error) {
	method := usage.NewMethod("serversettings").Add("max_players").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return 0, err
	}

	data, err := jsonrpc.DecodeRPCResult[int](r)
	if err != nil {
		return 0, err
	}

	return *data, nil
}

// SettingsMaxPlayersSet - Set the maximum number of players allowed to connect to the server
func (rpc *RPCClient) SettingsMaxPlayersSet(ctx context.Context, max int) (int, error) {
	method := usage.NewMethod("serversettings").Add("max_players").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, max)
	if err != nil {
		return 0, err
	}

	data, err := jsonrpc.DecodeRPCResult[int](r)
	if err != nil {
		return 0, err
	}

	return *data, nil
}

// ===========

func intSecDuration(secPtr *int) time.Duration {
	if secPtr == nil {
		return 0
	}
	return (time.Duration(*secPtr) * time.Second)
}

func durationIntSec(d time.Duration) int {
	return int(d.Seconds())
}

// SettingsPauseWhenEmptySeconds - Get the number of seconds before the game is automatically paused when no players are online
func (rpc *RPCClient) SettingsPauseWhenEmptySeconds(ctx context.Context) (time.Duration, error) {
	method := usage.NewMethod("serversettings").Add("pause_when_empty_seconds").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return 0, err
	}

	data, err := jsonrpc.DecodeRPCResult[int](r)
	if err != nil {
		return 0, err
	}

	return intSecDuration(data), nil
}

// SettingsPauseWhenEmptySecondsSet - Set the number of seconds before the game is automatically paused when no players are online
func (rpc *RPCClient) SettingsPauseWhenEmptySecondsSet(ctx context.Context, duration time.Duration) (time.Duration, error) {
	method := usage.NewMethod("serversettings").Add("pause_when_empty_seconds").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, durationIntSec(duration))
	if err != nil {
		return 0, err
	}

	data, err := jsonrpc.DecodeRPCResult[int](r)
	if err != nil {
		return 0, err
	}

	return intSecDuration(data), nil
}

// ===========

// SettingsPlayerIdleTimeout - Get the number of seconds before idle players are automatically kicked from the server
func (rpc *RPCClient) SettingsPlayerIdleTimeout(ctx context.Context) (time.Duration, error) {
	method := usage.NewMethod("serversettings").Add("player_idle_timeout").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return 0, err
	}

	data, err := jsonrpc.DecodeRPCResult[int](r)
	if err != nil {
		return 0, err
	}

	return intSecDuration(data), nil
}

// SettingsPlayerIdleTimeoutSet - Set the number of seconds before idle players are automatically kicked from the server
func (rpc *RPCClient) SettingsPlayerIdleTimeoutSet(ctx context.Context, duration time.Duration) (time.Duration, error) {
	method := usage.NewMethod("serversettings").Add("player_idle_timeout").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, durationIntSec(duration))
	if err != nil {
		return 0, err
	}

	data, err := jsonrpc.DecodeRPCResult[int](r)
	if err != nil {
		return 0, err
	}

	return intSecDuration(data), nil
}

// ===========

// SettingsAllowFlight - Get whether flight is allowed for players in Survival mode
func (rpc *RPCClient) SettingsAllowFlight(ctx context.Context) (bool, error) {
	method := usage.NewMethod("serversettings").Add("allow_flight").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return false, err
	}

	data, err := jsonrpc.DecodeRPCResult[bool](r)
	if err != nil {
		return false, err
	}

	return *data, nil
}

// SettingsAllowFlightSet - Set whether flight is allowed for players in Survival mode
func (rpc *RPCClient) SettingsAllowFlightSet(ctx context.Context, allow bool) (bool, error) {
	method := usage.NewMethod("serversettings").Add("allow_flight").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, allow)
	if err != nil {
		return false, err
	}

	data, err := jsonrpc.DecodeRPCResult[bool](r)
	if err != nil {
		return false, err
	}

	return *data, nil
}

// ===========

// SettingsMotd - Get the server's message of the day displayed to players
func (rpc *RPCClient) SettingsMotd(ctx context.Context) (string, error) {
	method := usage.NewMethod("serversettings").Add("motd").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return "", err
	}

	data, err := jsonrpc.DecodeRPCResult[string](r)
	if err != nil {
		return "", err
	}

	return *data, nil
}

// SettingsMotdSet - Set the server's message of the day displayed to players
func (rpc *RPCClient) SettingsMotdSet(ctx context.Context, motd string) (string, error) {
	method := usage.NewMethod("serversettings").Add("motd").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, motd)
	if err != nil {
		return "", err
	}

	data, err := jsonrpc.DecodeRPCResult[string](r)
	if err != nil {
		return "", err
	}

	return *data, nil
}

// ===========

// SettingsSpawnProtectionRadius - Get the spawn protection radius in blocks
func (rpc *RPCClient) SettingsSpawnProtectionRadius(ctx context.Context) (int, error) {
	method := usage.NewMethod("serversettings").Add("spawn_protection_radius").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return 0, err
	}

	data, err := jsonrpc.DecodeRPCResult[int](r)
	if err != nil {
		return 0, err
	}

	return *data, nil
}

// SettingsSpawnProtectionRadiusSet - Set the spawn protection radius in blocks
func (rpc *RPCClient) SettingsSpawnProtectionRadiusSet(ctx context.Context, radius int) (int, error) {
	method := usage.NewMethod("serversettings").Add("spawn_protection_radius").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, radius)
	if err != nil {
		return 0, err
	}

	data, err := jsonrpc.DecodeRPCResult[int](r)
	if err != nil {
		return 0, err
	}

	return *data, nil
}

// ===========

// SettingsForceGameMode - Get whether players are forced to use the server's default game mode
func (rpc *RPCClient) SettingsForceGameMode(ctx context.Context) (bool, error) {
	method := usage.NewMethod("serversettings").Add("force_game_mode").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return false, err
	}

	data, err := jsonrpc.DecodeRPCResult[bool](r)
	if err != nil {
		return false, err
	}

	return *data, nil
}

// SettingsForceGameModeSet - Set whether players are forced to use the server's default game mode
func (rpc *RPCClient) SettingsForceGameModeSet(ctx context.Context, forced bool) (bool, error) {
	method := usage.NewMethod("serversettings").Add("force_game_mode").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, forced)
	if err != nil {
		return false, err
	}

	data, err := jsonrpc.DecodeRPCResult[bool](r)
	if err != nil {
		return false, err
	}

	return *data, nil
}

// ===========

// SettingsGameMode - Get the server's default game mode
func (rpc *RPCClient) SettingsGameMode(ctx context.Context) (string, error) {
	method := usage.NewMethod("serversettings").Add("game_mode").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return "", err
	}

	data, err := jsonrpc.DecodeRPCResult[string](r)
	if err != nil {
		return "", err
	}

	return *data, nil
}

// SettingsGameModeSet - Set the server's default game mode
func (rpc *RPCClient) SettingsGameModeSet(ctx context.Context, gamemode string) (string, error) {
	method := usage.NewMethod("serversettings").Add("game_mode").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, gamemode)
	if err != nil {
		return "", err
	}

	data, err := jsonrpc.DecodeRPCResult[string](r)
	if err != nil {
		return "", err
	}

	return *data, nil
}

// ===========

// SettingsViewDistance - Get the server's view distance in chunks
func (rpc *RPCClient) SettingsViewDistance(ctx context.Context) (int, error) {
	method := usage.NewMethod("serversettings").Add("view_distance").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return 0, err
	}

	data, err := jsonrpc.DecodeRPCResult[int](r)
	if err != nil {
		return 0, err
	}

	return *data, nil
}

// SettingsViewDistanceSet - Set the server's view distance in chunks
func (rpc *RPCClient) SettingsViewDistanceSet(ctx context.Context, distance int) (int, error) {
	method := usage.NewMethod("serversettings").Add("view_distance").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, distance)
	if err != nil {
		return 0, err
	}

	data, err := jsonrpc.DecodeRPCResult[int](r)
	if err != nil {
		return 0, err
	}

	return *data, nil
}

// ===========

// SettingsSimulationDistance - Get the server's simulation distance in chunks
func (rpc *RPCClient) SettingsSimulationDistance(ctx context.Context) (int, error) {
	method := usage.NewMethod("serversettings").Add("simulation_distance").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return 0, err
	}

	data, err := jsonrpc.DecodeRPCResult[int](r)
	if err != nil {
		return 0, err
	}

	return *data, nil
}

// SettingsSimulationDistanceSet - Set the server's simulation distance in chunks
func (rpc *RPCClient) SettingsSimulationDistanceSet(ctx context.Context, distance int) (int, error) {
	method := usage.NewMethod("serversettings").Add("simulation_distance").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, distance)
	if err != nil {
		return 0, err
	}

	data, err := jsonrpc.DecodeRPCResult[int](r)
	if err != nil {
		return 0, err
	}

	return *data, nil
}

// ===========

// SettingsAcceptTransfers - Get whether the server accepts player transfers from other servers
func (rpc *RPCClient) SettingsAcceptTransfers(ctx context.Context) (bool, error) {
	method := usage.NewMethod("serversettings").Add("accept_transfers").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return false, err
	}

	data, err := jsonrpc.DecodeRPCResult[bool](r)
	if err != nil {
		return false, err
	}

	return *data, nil
}

// SettingsAcceptTransfersSet - Set whether the server accepts player transfers from other servers
func (rpc *RPCClient) SettingsAcceptTransfersSet(ctx context.Context, accept bool) (bool, error) {
	method := usage.NewMethod("serversettings").Add("accept_transfers").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, accept)
	if err != nil {
		return false, err
	}

	data, err := jsonrpc.DecodeRPCResult[bool](r)
	if err != nil {
		return false, err
	}

	return *data, nil
}

// ===========

// SettingsStatusHeartbeatInterval - Get the interval in seconds between server status heartbeats
func (rpc *RPCClient) SettingsStatusHeartbeatInterval(ctx context.Context) (time.Duration, error) {
	method := usage.NewMethod("serversettings").Add("status_heartbeat_interval").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return 0, err
	}

	data, err := jsonrpc.DecodeRPCResult[int](r)
	if err != nil {
		return 0, err
	}

	return intSecDuration(data), nil
}

// SettingsStatusHeartbeatIntervalSet - Set the interval in seconds between server status heartbeats
func (rpc *RPCClient) SettingsStatusHeartbeatIntervalSet(ctx context.Context, interval time.Duration) (time.Duration, error) {
	method := usage.NewMethod("serversettings").Add("status_heartbeat_interval").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, durationIntSec(interval))
	if err != nil {
		return 0, err
	}

	data, err := jsonrpc.DecodeRPCResult[int](r)
	if err != nil {
		return 0, err
	}

	return intSecDuration(data), nil
}

// ===========

// SettingsOperatorUserPermissionLevel - Get the permission level required for operator commands
func (rpc *RPCClient) SettingsOperatorUserPermissionLevel(ctx context.Context) (int, error) {
	method := usage.NewMethod("serversettings").Add("operator_user_permission_level").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return 0, err
	}

	data, err := jsonrpc.DecodeRPCResult[int](r)
	if err != nil {
		return 0, err
	}

	return *data, nil
}

// SettingsOperatorUserPermissionLevelSet - Set the permission level required for operator commands
func (rpc *RPCClient) SettingsOperatorUserPermissionLevelSet(ctx context.Context, level int) (int, error) {
	method := usage.NewMethod("serversettings").Add("operator_user_permission_level").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, level)
	if err != nil {
		return 0, err
	}

	data, err := jsonrpc.DecodeRPCResult[int](r)
	if err != nil {
		return 0, err
	}

	return *data, nil
}

// ===========

// SettingsHideOnlinePlayers - Get whether the server hides online player information from status queries
func (rpc *RPCClient) SettingsHideOnlinePlayers(ctx context.Context) (bool, error) {
	method := usage.NewMethod("serversettings").Add("hide_online_players").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return false, err
	}

	data, err := jsonrpc.DecodeRPCResult[bool](r)
	if err != nil {
		return false, err
	}

	return *data, nil
}

// SettingsHideOnlinePlayersSet - Set whether the server hides online player information from status queries
func (rpc *RPCClient) SettingsHideOnlinePlayersSet(ctx context.Context, hide bool) (bool, error) {
	method := usage.NewMethod("serversettings").Add("hide_online_players").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, hide)
	if err != nil {
		return false, err
	}

	data, err := jsonrpc.DecodeRPCResult[bool](r)
	if err != nil {
		return false, err
	}

	return *data, nil
}

// ===========

// SettingsStatusReplies - Get whether the server responds to connection status requests
func (rpc *RPCClient) SettingsStatusReplies(ctx context.Context) (bool, error) {
	method := usage.NewMethod("serversettings").Add("status_replies").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return false, err
	}

	data, err := jsonrpc.DecodeRPCResult[bool](r)
	if err != nil {
		return false, err
	}

	return *data, nil
}

// SettingsStatusRepliesSet - Set whether the server responds to connection status requests
func (rpc *RPCClient) SettingsStatusRepliesSet(ctx context.Context, enabled bool) (bool, error) {
	method := usage.NewMethod("serversettings").Add("status_replies").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, enabled)
	if err != nil {
		return false, err
	}

	data, err := jsonrpc.DecodeRPCResult[bool](r)
	if err != nil {
		return false, err
	}

	return *data, nil
}

// ===========

// SettingsEntityBroadcastRange - Get the entity broadcast range as a percentage
func (rpc *RPCClient) SettingsEntityBroadcastRange(ctx context.Context) (int, error) {
	method := usage.NewMethod("serversettings").Add("entity_broadcast_range").String()
	r, err := rpc.core.CallWithContext(ctx, method)
	if err != nil {
		return 0, err
	}

	data, err := jsonrpc.DecodeRPCResult[int](r)
	if err != nil {
		return 0, err
	}

	return *data, nil
}

// SettingsEntityBroadcastRangeSet - Set the entity broadcast range as a percentage
func (rpc *RPCClient) SettingsEntityBroadcastRangeSet(ctx context.Context, percentage_points int) (int, error) {
	method := usage.NewMethod("serversettings").Add("entity_broadcast_range").Add("set").String()
	r, err := rpc.core.CallWithContext(ctx, method, percentage_points)
	if err != nil {
		return 0, err
	}

	data, err := jsonrpc.DecodeRPCResult[int](r)
	if err != nil {
		return 0, err
	}

	return *data, nil
}
