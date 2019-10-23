# Plugins

This document aims to plan and describe the plugin framework for _coypu_.

Coypu will (probably) ship with some basic plugins, enough to run.

There are different types of plugins, some manage input (such as database or http proxy), some manage output (such as yaml, json, or html table), some act as middleware (such as filtering out records that don't match validation), and some manage authorization (such as jwt, external url).

# Plugin Interfaces

The following endpoints are expected to exist for each type:

Auth Plugins should have an interface like:
`success = auth(\*context)`

Input Plugins should have an interface like:
`input_result = input(\*context, query, http_method, body)`

Modification Plugins should have an interface like:
`mod_result = modify(\*context, input_result)`

Output Plugins should have an interface like:
`success = render(\*context, input_or_mod_result)``

The following types:
map: context, input_result, mod_result, query
boolean: success
string: http_method

All interfaces can raise errors, which will be reported to the user. Do not put anything sensitive in error messages.

TODO *it's probably a good idea to semi-formalize context map too (standard fields)*
