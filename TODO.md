* Make project structure not flat?
  - Could break everything into packages - models, handlers, store, etc

* Add jwt auth for resources and _only_ return resource for proper user
  - Currently working on doing this for work blocks

* Tests.

* Add more necessary db indexes.

* Account for cascading?

* Should projects belong to users too?

* Handle nil values being sent to update actions

* UUID and "sessions"/jwt for users.

* UUID for all id values.

* Dates
  - Should the struct field be something else? Or does this save a lot of trouble?

* Shared functions for similar actions in stores and handlers
  - Especially pay period handlers

* Can an array of struct pointers be passed into index handlers, so that it is more
on par with other actions?
  - PayPeriodsIndexHandler
    1. Should ID be part of an object instead of an argument?
