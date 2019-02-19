# TODO

* Add colored text for sub-fields in record (w/ custom directive)
* Add a timeline of events

|     |           |   |                  |           |          |                                                                 |
|-----|-----------|---|------------------|-----------|----------|-----------------------------------------------------------------|
| The | PROJECT   | X | has been updated | ${Before} | ${After} | ONCE AT CREATION                                                |
| The | DATASET   | X | has been updated | ${Before} | ${After} | ONCE AT CREATION / EACH PERMISSION CHANGES                      |
| The | TABLE     | X | has been updated | ${Before} | ${After} | ONCE AT CREATION / EACH PERMISSION CHANGES                      |
| The | PARTITION | X | has been updated | ${Before} | ${After} | ONCE AT CREATION / EACH COUNT CHANGES                           |
| The | FIELD     | X | has been updated | ${Before} | ${After} | ONCE AT CREATION / EACH DESCRIPTION CHANGES / EACH MODE CHANGES |

* Add some graph widgets to make the app sexier
* Add colors in tables to make the app sexier
* Handle BigQuery job errors
* Filter by label and not by dataset name
* Add a graph based on level project -> dataset -> table -> field
* Add a graph per table with parent tables and child tables
* New icon