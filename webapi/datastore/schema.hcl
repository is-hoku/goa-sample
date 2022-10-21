table "students" {
  schema = schema.temp
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    auto_increment = true
  }
  column "name" {
    null = false
    type = varchar(128)
  }
  column "ruby" {
    null = false
    type = varchar(128)
  }
  column "student_number" {
    null     = false
    type     = int
    unsigned = true
  }
  column "date_of_birth" {
    null = false
    type = datetime
  }
  column "address" {
    null = false
    type = varchar(256)
  }
  column "expiration_date" {
    null = false
    type = datetime
  }
  column "created_at" {
    null    = false
    type    = datetime
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null      = false
    type      = datetime
    default   = sql("CURRENT_TIMESTAMP")
    on_update = sql("CURRENT_TIMESTAMP")
  }
  primary_key {
    columns = [column.id]
  }
  index "student_number" {
    unique  = true
    columns = [column.student_number]
  }
}
schema "temp" {
  charset = "utf8mb4"
  collate = "utf8mb4_bin"
}
