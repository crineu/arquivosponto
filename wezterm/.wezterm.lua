-- Pull in the wezterm API
local wezterm = require("wezterm")

-- This will hold the configuration.
local config = wezterm.config_builder()

-- This is where you actually apply your config choices

-- For example, changing the color scheme:
-- config.color_scheme = 'Tokyo Night'
-- config.color_scheme = 'GruvboxDark'
-- config.color_scheme = 'Catppuccin Frappe'
-- config.color_scheme = 'nord'
config.color_scheme = "Kanagawa (Gogh)"
-- config.color_scheme = 'rose-pine'

config.initial_cols = 152
config.initial_rows = 38

config.font_size = 13
wezterm.font_with_fallback({
	-- <built-in>, BuiltIn
	"JetBrains Mono",

	-- <built-in>, BuiltIn
	-- Assumed to have Emoji Presentation
	"Noto Color Emoji",

	-- <built-in>, BuiltIn
	"Symbols Nerd Font Mono",
})

config.tab_bar_at_bottom = true
config.keys = {
	{
		key = "LeftArrow",
		mods = "SUPER",
		action = wezterm.action.ActivateTabRelative(-1),
	},
	{
		key = "RightArrow",
		mods = "SUPER",
		action = wezterm.action.ActivateTabRelative(1),
	},
	{
		key = ",",
		mods = "SUPER",
		action = wezterm.action.PromptInputLine({
			description = "New name for tab",
			action = wezterm.action_callback(function(window, pane, line)
				if line then
					window:active_tab():set_title(line)
				end
			end),
		}),
	},
}

-- and finally, return the configuration to wezterm
return config
