* REMEMBER: Hours should be calculated by start/end time for a work period, and
then returned. Not received as params

* Tests. I really should test this time, at the very least, for the learning experience.

* Validations for pay period started_on being before ended_on
  - Do this on the db level?

* Handle nil values being sent to update actions

* UUID and "sessions"/jwt for users.

* UUID for all id values.

* Date formatting - should it return as "03-09-1990" date to client, or with timezone?
  - It could probably take in timezone as well. So maybe it's all on the client.
  - Should the struct field be something else? Or does this save a lot of trouble?

* Shared functions for similar actions in stores and handlers
  - Especially pay period handlers

* Can an array of struct pointers be passed into index handlers, so that it is more
on par with other actions?
  - PayPeriodsIndexHandler
    1. Should ID be part of an object instead of an argument?

* Should projects be limited by user, or user group maybe?