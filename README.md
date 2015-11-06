# EBMF
Event-Based Monitoring Framework

# Discussed entry points

## General
* We want an event-based monitoring framework
* We want to de-couple monitoring data collection and alerting logic
* We want to have a distributed and resilient system that's not tied to any single points of failure
* We want a way to run monitoring plugins and collect their output
* We want a push-based system
* We want a defined protocol for the monitoring data collection with a possibility to extend it with user-defined data
* We want to remain Nagios-compatible to be able to "drop-in" enhance existing monitoring systems with the possibility of replacing them
* We do not want any proprietary exchange protocols. In other words - client-server communication should use HTTP with JSON as data exchange format
* We do not want to vendor-lock instances to our agent. We want to allow applications to report monitoring data to the server without the need to write plugins outside of the application (think Codahale and Graphite)
* We want "Graphite for non-integer data"

## Agent side
* Agent is stateless
* Agent collects monitoring "check data" on the instance/node and passes it to the server without making any decisions on the state
* Agent only connects and talks to the server on its own, does not allow for any data to be pushed to it
* Agent defines what will be monitored, how and at which periods and communicates this information to server upon startup and/or upon configuration changes
* Configuration of the agent can possibly be stored in consul or be otherwise distributed as a static file with it that's re-read frequently

## Server side
* Server registers expected monitoring data (which datapoints will arrive and at which periods), based on the registration call from the agent
* Server stores configuration data received from agent until instructed to clean it up by the agent or via an API call
* Server passively gathers data supplied by the agents
* Server makes sure that expected data from agents is coming in at expected intervals, otherwise assumes (part of the) agent dead
* Server stores data received from the agents "as-is" (JSON blobs)
* Server exposes an API to access stored data, search through it, aggregate it, do processing (TBD)
* Server exposes an API to define cron-like jobs to process stored data
* Server exposes an API to define actions (send e-mail, send slack message, send sms, smack the user, etc)
* Server exposes an API to couple jobs to actions based on user-defined conditions

## UI
* Excplicitly based on server-API, no direct data access
* UI is a domain where I'm a little bit braindead, but will give it some thought later on, otherwise feel free to fill this out


