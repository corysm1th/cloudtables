# CloudTables

```plantuml
@startuml

!includeurl https://raw.githubusercontent.com/corysm1th/script_dumpster/master/style.puml

participant "User \nInterface" as ui
participant "CloudTables \nAPI" as srv
participant "func Sync()" as sync
database "Database" as db
participant "Cloud \nProvider" as cloud

note right of srv
Read config
    Addr
    CertFile
    KeyFile
    CAFile
    MutualAuth
if MutualAuth {
    Run ServeMutualAuth()
else
    Run ServeTLS()
end note

note right of srv
Read accounts
    Account Name
    User Name
end note

group For each account
srv -> db : store account in accounts
note right of db
Account
    Account name
    User name
    Sync state
    Last sync'd
end note
end

srv -> sync : sync

group For each account
sync -> sync : if Account.SyncState = InProgress {return}
sync -> db : Account.SyncState = InProgress
sync -> cloud : Get all objects
group if err != nil
sync -> db : Account.SyncState = Error
sync -> sync : log err
end
sync -> db : Store all objects
sync -> db : Account.SycState = Complete
sync -> sync : Calculate metrics
sync -> db : store metrics
end

ui -> srv : GET /api/v1/objects
srv -> db : select all from all

note right of db
Objects
    AWS
    GCP
    Azure
end note

srv <- db : return all
ui <- srv : return 200, json payload
ui -> srv : GET /api/v1/sync
srv -> sync : sync
ui <- srv : return 201
note right of ui
sync button grey'd out 1 min
"Account sync in progress...
Check Admin page for details
end note
ui -> srv : GET /api/v1/metrics
srv -> db : get metrics
srv <- db : metrics
ui <- srv : return 200, json payload
@enduml
```