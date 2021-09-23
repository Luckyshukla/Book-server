package repository
import(
	//"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"server.go/entity"
)
type VideoRepository interface{
	Save(video entity.Video)
	Updated()
	Delete()
	FindAll() []entity.Video
	CloseDB()
}
type database struct{
	connection *gorm.DB
}
func NewVideoRepository() VideoRepository {
	//db, err := gorm.Open(sqlite.Open("test.db")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.Connection.Cose()
	if err !=nil{
		panic("Faild to close database")
	}
}
func (db *database) Save(video entity.Video) {
	db.connection.Create(&video)
}

func (db *database) Update(video entity.Video) {
	db.connection.Save(&video)
}

func (db *database) Delete(video entity.Video) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}
