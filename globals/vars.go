package globals

import "gorm.io/gorm"

var Db *gorm.DB

// I know, so funny, right? Please laugh
var JwtSecret = []byte("i am very hard to crack")
