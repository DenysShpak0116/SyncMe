package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"server/dto"
	"server/models"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Health() map[string]string
	Close() error

	AddUser(user models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByUsername(login string) (*models.User, error)
	GetUserById(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	ChangeUserRole(userId int, role string) error

	AddGroup(group models.Group) (int, error)
	GetAllGroups() []models.Group
	GetGroupById(id int) (*models.Group, error)

	AddUserGroup(userId int, groupId int) error

	AddAuthor(author models.Author) (int, error)
	GetAuthorById(id int) (*models.Author, error)
	GetAuthorsByGroupId(groupId int) ([]models.Author, error)

	AddPost(post models.Post) (int, error)
	GetPostsByAuthorId(authorId int) ([]models.Post, error)

	AddPhoto(photo models.XPhoto) (int, error)
	GetPhotosByPostId(postId int) ([]models.XPhoto, error)

	AddVideo(video models.XVideo) (int, error)
	GetVideosByPostId(postId int) ([]models.XVideo, error)

	AddMessage(message models.Message) (int, error)

	DeleteMessage(messageId int) error
	GetChat(disscusserId int, currentUserId int) ([]models.Message, error)

	AddComment(comment models.Comment) (int, error)
	GetPostComments(postId int) ([]models.Comment, error)

	GetUserChats(userId int) ([]dto.Chat, error)

	AddEmotionalAnalysis(emotionalAnalysis models.EmotionalAnalysis) (int, error)
	GetEmotionalAnalysisById(id int) (*models.EmotionalAnalysis, error)
	GetAuthorEmotionalAnalysis(authorId int) (*dto.EmotionalAnalysis, error)
	GetGroupEmotionalAnalysis(groupId int) (*dto.EmotionalAnalysis, error)

	AddNotification(notificationRequest models.Notification) (int, error)
	AddUserNotification(userId int, notificationId int) (int, error)
	GetUserNotifications(userId int) ([]models.Notification, error)
}

type service struct {
	db *sql.DB
}

var (
	dbname     = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	dbInstance *service
)

func Instance() Service {
	return dbInstance
}

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	// Opening a driver typically will not attempt to connect to the database.
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			username,
			password,
			host,
			port,
			dbname,
		),
	)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	dbInstance = &service{
		db: db,
	}

	log.Println("Connected")
	return dbInstance
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf(fmt.Sprintf("db down: %v", err)) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}
	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dbname)
	return s.db.Close()
}

func (s *service) AddUser(user models.User) error {
	query := `INSERT INTO user (username, password, email, firstname, lastname, sex, country, role) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Sex,
		user.Country,
		user.Role,
	)
	if err != nil {
		return fmt.Errorf("could not insert user: %v", err)
	}

	return nil
}

func (s *service) GetAllUsers() ([]models.User, error) {
	query := `SELECT * FROM user`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve users: %v", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.UserId,
			&user.Username,
			&user.Password,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Sex,
			&user.Country,
			&user.Role,
			&user.Logo,
			&user.BgImage,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan user: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over users: %v", err)
	}

	return users, nil
}

func (s *service) GetUserByUsername(username string) (*models.User, error) {
	query := `SELECT UserId, Username, Password, Email, FirstName,LastName, Sex, Country, Role, Logo, BgImage FROM user WHERE Username = ?`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	row := s.db.QueryRowContext(ctx, query, username)
	var user models.User
	err := row.Scan(
		&user.UserId,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Sex,
		&user.Country,
		&user.Role,
		&user.Logo,
		&user.BgImage,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (s *service) GetUserById(id int) (*models.User, error) {
	query := `SELECT UserId, Username, Password, Email, FirstName, LastName, Sex, Country, Role, Logo, BgImage FROM user WHERE UserId = ?`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	row := s.db.QueryRowContext(ctx, query, id)
	var user models.User
	err := row.Scan(
		&user.UserId,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Sex,
		&user.Country,
		&user.Role,
		&user.Logo,
		&user.BgImage,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *service) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT UserId, Username, Password, Email, FirstName, LastName, Sex, Country, Role, Logo, BgImage FROM user WHERE Email = ?`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	row := s.db.QueryRowContext(ctx, query, email)
	var user models.User
	err := row.Scan(
		&user.UserId,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Sex,
		&user.Country,
		&user.Role,
		&user.Logo,
		&user.BgImage,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *service) AddGroup(group models.Group) (int, error) {
	query := "INSERT INTO `group` (Name, GroupImage, GroupBackgroundImage, Description, EmotionalAnalysisId) VALUES (?, ?, ?, ?, ?)"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := s.db.ExecContext(ctx, query, group.Name, group.GroupImage, group.GroupBackgroundImage, group.Description, 1)
	if err != nil {
		return -1, fmt.Errorf("could not insert group: %v", err.Error())
	}
	groupId, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("could not retrieve group id: %v", err.Error())
	}
	return int(groupId), nil
}

func (s *service) AddUserGroup(userId int, groupId int) error {
	query := "INSERT INTO usergroup (UserId, GroupId) VALUES (?, ?)"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userId, groupId)
	if err != nil {
		return fmt.Errorf("could not insert user group: %v", err.Error())
	}
	return nil
}

func (s *service) AddAuthor(author models.Author) (int, error) {
	query := `INSERT INTO author (Name, Username, SocialMedia, AuthorImage, 
		AuthorBackgroundImage, GroupId, EmotionalAnalysisId) VALUES (?, ?, ?, ?, ?, ?, ?)`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := s.db.ExecContext(
		ctx,
		query,
		author.Name,
		author.Username,
		author.SocialMedia,
		author.AuthorImage,
		author.AuthorBackgroundImage,
		author.GroupId,
		1,
	)
	if err != nil {
		return -1, fmt.Errorf("could not insert author: %v", err.Error())
	}
	authorId, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("could not retrieve author id: %v", err.Error())
	}
	return int(authorId), nil
}

func (s *service) GetAllGroups() []models.Group {
	query := "SELECT * FROM `group`"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := dbInstance.db.QueryContext(ctx, query)
	if err != nil {
		log.Fatalf("could not retrieve groups: %v", err)
	}
	defer rows.Close()
	var groups []models.Group
	for rows.Next() {
		var group models.Group
		err := rows.Scan(
			&group.GroupId,
			&group.Name,
			&group.GroupImage,
			&group.GroupBackgroundImage,
			&group.EmotionalAnalysisId,
			&group.Description,
		)
		if err != nil {
			log.Fatalf("could not scan group: %v", err)
		}
		groups = append(groups, group)
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("error iterating over groups: %v", err)
	}
	return groups
}

func (s *service) GetGroupById(id int) (*models.Group, error) {
	query := "SELECT GroupId, Name, GroupImage, GroupBackgroundImage, Description, EmotionalAnalysisId FROM `group` WHERE GroupId = ?"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	row := dbInstance.db.QueryRowContext(ctx, query, id)
	var group models.Group
	err := row.Scan(
		&group.GroupId,
		&group.Name,
		&group.GroupImage,
		&group.GroupBackgroundImage,
		&group.Description,
		&group.EmotionalAnalysisId,
	)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (s *service) GetAuthorsByGroupId(groupId int) ([]models.Author, error) {
	query := "SELECT AuthorId, Name, Username, SocialMedia, AuthorImage, AuthorBackgroundImage, GroupId, EmotionalAnalysisId FROM author WHERE GroupId = ?"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := dbInstance.db.QueryContext(ctx, query, groupId)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve authors: %v", err)
	}
	defer rows.Close()
	var authors []models.Author
	for rows.Next() {
		var author models.Author
		err := rows.Scan(
			&author.AuthorId,
			&author.Name,
			&author.Username,
			&author.SocialMedia,
			&author.AuthorImage,
			&author.AuthorBackgroundImage,
			&author.GroupId,
			&author.EmotionalAnalysisId,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan author: %v", err)
		}
		authors = append(authors, author)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over authors: %v", err)
	}
	return authors, nil
}

func (s *service) GetAuthorById(id int) (*models.Author, error) {
	query := "SELECT AuthorId, Name, Username, SocialMedia, AuthorImage, AuthorBackgroundImage, GroupId, EmotionalAnalysisId FROM author WHERE AuthorId = ?"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	row := dbInstance.db.QueryRowContext(ctx, query, id)
	var author models.Author
	err := row.Scan(
		&author.AuthorId,
		&author.Name,
		&author.Username,
		&author.SocialMedia,
		&author.AuthorImage,
		&author.AuthorBackgroundImage,
		&author.GroupId,
		&author.EmotionalAnalysisId,
	)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (s *service) AddPost(post models.Post) (int, error) {
	query := `INSERT INTO post (TextContent, Date, CountOfLikes, AuthorId, EmotionalAnalysisId) VALUES (?, ?, ?, ?, ?)`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := s.db.ExecContext(
		ctx,
		query,
		post.TextContent,
		post.Date,
		post.CountOfLikes,
		post.AuthorId,
		post.EmotionalAnalysisId,
	)
	if err != nil {
		return -1, fmt.Errorf("could not insert post: %v", err)
	}
	postId, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("could not retrieve post id: %v", err)
	}
	return int(postId), nil
}

func (s *service) AddPhoto(photo models.XPhoto) (int, error) {
	query := `INSERT INTO xphoto (URL, PostId) VALUES (?, ?)`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := s.db.ExecContext(
		ctx,
		query,
		photo.URL,
		photo.PostId,
	)
	if err != nil {
		return -1, fmt.Errorf("could not insert photo: %v", err.Error())
	}
	photoId, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("could not retrieve photo id: %v", err.Error())
	}
	return int(photoId), nil
}

func (s *service) AddVideo(video models.XVideo) (int, error) {
	query := `INSERT INTO xvideo (URL, PostId) VALUES (?, ?)`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := s.db.ExecContext(
		ctx,
		query,
		video.URL,
		video.PostId,
	)
	if err != nil {
		return -1, fmt.Errorf("could not insert video: %v", err.Error())
	}
	videoId, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("could not retrieve video id: %v", err.Error())
	}
	return int(videoId), nil
}

func (s *service) GetPostsByAuthorId(authorId int) ([]models.Post, error) {
	query := `SELECT PostId, TextContent, Date, CountOfLikes, AuthorId, EmotionalAnalysisId FROM post WHERE AuthorId = ?`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := s.db.QueryContext(ctx, query, authorId)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve posts: %v", err)
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.PostId,
			&post.TextContent,
			&post.Date,
			&post.CountOfLikes,
			&post.AuthorId,
			&post.EmotionalAnalysisId,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan post: %v", err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over posts: %v", err)
	}
	return posts, nil
}

func (s *service) GetPhotosByPostId(postId int) ([]models.XPhoto, error) {
	query := `SELECT XPhotoId, URL, PostId FROM xphoto WHERE PostId = ?`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := s.db.QueryContext(ctx, query, postId)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve photos: %v", err)
	}
	defer rows.Close()
	var photos []models.XPhoto
	for rows.Next() {
		var photo models.XPhoto
		err := rows.Scan(
			&photo.XPhotoId,
			&photo.URL,
			&photo.PostId,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan photo: %v", err)
		}
		photos = append(photos, photo)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over photos: %v", err)
	}
	return photos, nil
}

func (s *service) GetVideosByPostId(postId int) ([]models.XVideo, error) {
	query := `SELECT XVideoId, URL, PostId FROM xvideo WHERE PostId = ?`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := s.db.QueryContext(ctx, query, postId)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve videos: %v", err)
	}
	defer rows.Close()
	var videos []models.XVideo
	for rows.Next() {
		var video models.XVideo
		err := rows.Scan(
			&video.XVideoId,
			&video.URL,
			&video.PostId,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan video: %v", err)
		}
		videos = append(videos, video)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over videos: %v", err)
	}
	return videos, nil
}

func (s *service) AddMessage(message models.Message) (int, error) {
	query := `INSERT INTO message (Text, SentAt, UserFromId, UserToId) VALUES (?, ?, ?, ?)`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := s.db.ExecContext(
		ctx,
		query,
		message.Text,
		message.SentAt,
		message.UserFromId,
		message.UserToId,
	)
	if err != nil {
		return -1, fmt.Errorf("could not insert message: %v", err)
	}
	messageId, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("could not retrieve message id: %v", err)
	}
	return int(messageId), nil
}

func (s *service) GetChat(disscusserId int, currentUserId int) ([]models.Message, error) {
	query := `SELECT MessageId, Text, SentAt, UserFromId, UserToId FROM message WHERE (UserFromId = ? AND UserToId = ?) OR (UserFromId = ? AND UserToId = ?) ORDER BY SentAt`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := s.db.QueryContext(ctx, query, disscusserId, currentUserId, currentUserId, disscusserId)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve messages: %v", err)
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var message models.Message
		err := rows.Scan(
			&message.MessageId,
			&message.Text,
			&message.SentAt,
			&message.UserFromId,
			&message.UserToId,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan message: %v", err)
		}
		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over messages: %v", err)
	}
	return messages, nil
}

func (s *service) DeleteMessage(messageId int) error {
	query := `DELETE FROM message WHERE MessageId = ?`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, messageId)
	if err != nil {
		return fmt.Errorf("could not delete message: %v", err)
	}
	return nil
}

func (s *service) AddComment(comment models.Comment) (int, error) {
	query := `INSERT INTO comment (Text, Date, UserId, PostId) VALUES (?, ?, ?, ?)`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := s.db.ExecContext(
		ctx,
		query,
		comment.Text,
		comment.Date,
		comment.UserId,
		comment.PostId,
	)
	if err != nil {
		return -1, fmt.Errorf("could not insert comment: %v", err)
	}
	commentId, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("could not retrieve comment id: %v", err)
	}
	return int(commentId), nil
}

func (s *service) GetPostComments(postId int) ([]models.Comment, error) {
	query := `SELECT CommentId, Text, Date, UserId, PostId FROM comment WHERE PostId = ?`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := s.db.QueryContext(ctx, query, postId)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve comments: %v", err)
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(
			&comment.CommentId,
			&comment.Text,
			&comment.Date,
			&comment.UserId,
			&comment.PostId,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan comment: %v", err)
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over comments: %v", err)
	}
	return comments, nil
}

func (s *service) GetUserChats(userId int) ([]dto.Chat, error) {
	query := `SELECT DISTINCT UserFromId, UserToId FROM message WHERE UserFromId = ? OR UserToId = ?`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := s.db.QueryContext(ctx, query, userId, userId)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve chats: %v", err)
	}
	defer rows.Close()

	var chats []dto.Chat
	chatMap := make(map[int]bool) // Map to track unique chats
	for rows.Next() {
		var chat dto.Chat
		var userFromId, userToId int
		err := rows.Scan(
			&userFromId,
			&userToId,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan chat: %v", err)
		}
		// Check if the chat already exists in the map
		if _, ok := chatMap[userFromId+userToId]; ok {
			continue // Skip if chat already exists
		}
		chatMap[userFromId+userToId] = true // Add chat to the map
		if userFromId == userId {
			chat.UserId = userToId
		} else {
			chat.UserId = userFromId
		}
		chat.UserName, err = s.getUserName(chat.UserId)
		if err != nil {
			return nil, fmt.Errorf("could not get username: %v", err)
		}
		chats = append(chats, chat)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over chats: %v", err)
	}
	return chats, nil
}

func (s *service) getUserName(userId int) (string, error) {
	query := `SELECT Username FROM user WHERE UserId = ?`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	row := s.db.QueryRowContext(ctx, query, userId)
	var username string
	err := row.Scan(&username)
	if err != nil {
		return "", fmt.Errorf("could not get username: %v", err)
	}
	return username, nil
}

func (s *service) ChangeUserRole(userId int, role string) error {
	query := `UPDATE user SET Role = ? WHERE UserId = ?`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, role, userId)
	if err != nil {
		return fmt.Errorf("could not update user role: %v", err)
	}
	return nil
}
