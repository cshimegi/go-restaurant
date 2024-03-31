# Shared Models
Common models are shared among other services

## Model Declaration
### Ver1
```go
type BaseModel struct {
    CreatedAt time.Time `json:"created_at" gorm:"->:false;<-:create;not null"`
    UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}
```
This will make the created_at in response data json be
```json
{
  "id": 1,
  "first_name": "Alan",
  "last_name": "Chan",
  "nickname": "Alan Chan",
  "email": "xxx@hotmail.com",
  "password": "testtest",
  "created_at": "0001-01-01T00:00:00Z",
  "updated_at": "2022-07-25T12:34:56Z"
}
```
It replaces the real value of **created_at** with dummy value.

### Ver2
```go
type BaseModel struct {
    CreatedAt time.Time `json:"created_at" gorm:"->;<-:create;not null"`
    UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}
```
This will make the created_at in response data json be the same value in DB record
```json
{
  "id": 1,
  "first_name": "Alan",
  "last_name": "Chan",
  "nickname": "Alan Chan",
  "email": "xxx@hotmail.com",
  "password": "testtest",
  "created_at": "2022-07-25T12:34:56Z",
  "updated_at": "2022-07-25T12:34:56Z"
}
```



## Reference
[official gorm model declaration](https://gorm.io/docs/models.html#Field-Level-Permission)
