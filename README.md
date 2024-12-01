# Queudo
This projek is fo fun. Plz don't asq.

## Build
`go run .`

## Tech Stack
  - Golang
  - Sqlite
  - Bulma
  - Htmx
  - Alpine

The stack's objective is to allow quick development by removing the front-end build by using tools such as bulma, htmx, alpine and the backend setup of a full database server. I chose Golang just because I gotta try it serously at some point.

## TODO
### [ ] Phase 1: The Queue(tm)
In this phase I can add and clear my todos comming from different channels. We need to focus on the fact that this app should help me focus on 1 to 3 tasks at a single moment to get them done.

  - [ ] I can add channels and todos for those channels
  - [ ] I can view and clear todos of different channels from a single queue
  - [ ] Empty channels ask me to add something to them or disable them, because a channel should never be empty.
  - [ ] I have different channel grouping views (workspace)
  - [ ] I can put a time-out with a delay on my todos to temporarly disable them (usefull for things that has some waiting)
  - [ ] I can make a todo with at a specific moment (date/time) with some form of count down reminder (reminder offset should be defined for that todo)
  - [ ] I can make some todos recurring

### [ ] Phase 2: The Schedule
Todos can form a schedule to fit within a day based on UNLOAD (Uniquely Named Léa's Aesthetic Organizational Device).

### [ ] Phase 3: The Data
Todos can be more than a simple status, they can be complete forms that may require some complex information structure that are defined by the user.

### [ ] Phase 4: The API
We can query and mutate the app's data using an API.
It would be nice to be able to recover the data for use in BI tools.
