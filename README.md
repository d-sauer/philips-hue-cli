# Philips Hue cli - hue

Small CLI for interacting with Philips Hue lights.

## Commands

`hue discover` - discover Philips Hue hub locally 

`hue adduser` - add default user (`hue-cli`) to the Philips Hue Hub for this application, return `apikey` 

`hue adduser <name>` - add new user by `name` to the Philips Hue Hub for the application, return `apikey` 

`HUE_CLI_APIKEY` Environment variable for the API key

`hue lights` - get all available lights

`hue groups` - get all available lights

`hue on/off/dim` - manage lights state
 
- available flags
    - `-g` list of groups to apply command to
    - `-l` list of lights to apply command to
    - `-t` time interval in which to apply command
    - `-b` brightness or interval
    - `-s` saturation or interval
    - `-f` function: `linear`, `log`

