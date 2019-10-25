# Plugins

This document aims to plan and describe the plugin framework for _coypu_.

Coypu will (probably) ship with some basic plugins, enough to run.

# Plugin Interfaces

Plugins can be called in any order. They should each have a function called "New" which takes in the config object described in the plugin yaml document, and a returns a function which takes in and returns a context map (the plugin itself).

## Context Maps

The context map includes the http *method*, the *query* as a map, the request *body*, request *headers* as an array, *error*, *resStatus*, *resBody*, *resHeaders* for response sending, and *authDenied* as intended for an auth plugin to write to as needed.

At the end of all plugins, a response message with either a 5XX code and body *error* is sent, or a response message with a *resStatus* code, with body *resBody* and headers *resHeaders*.

Of course, plugins can write anything they would like to the context that they output to be used by other plugins in the stack. It's highly recommended to document this.
