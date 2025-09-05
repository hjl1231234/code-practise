package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;unique;not null"`
	Posts     []Post `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	PostCount int `gorm:"default:0"` // 新增文章数量字段
}

type Post struct {
	ID        uint      `gorm:"primarykey"`
	Title     string    `gorm:"size:200;not null"`
	Content   string    `gorm:"type:text"`
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
	Comments  []Comment `gorm:"foreignKey:PostID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CommentStatus string    `gorm:"size:20;default:'有评论'"` // 新增评论状态字段
}

type Comment struct {
	ID        uint   `gorm:"primarykey"`
	Content   string `gorm:"type:text;not null"`
	PostID    uint   `gorm:"not null"`
	Post      Post   `gorm:"foreignKey:PostID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	dsn := "host=localhost user=postgres password=pwd dbname=test_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: false})
	if err != nil {
		log.Fatal(err)
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("数据库表创建成功")

	// 如果数据库没有数据，生成一些测试数据
	var userCount int64
	db.Model(&User{}).Count(&userCount)
	if userCount == 0 {
		InsertRandomData(db)
	}

	// test_createPost(db)
	// test_CreateComment(db)

	// 调用查询函数并打印结果
	// 1. 查询第一个用户的所有文章及其评论
	var firstUser User
	db.First(&firstUser)
	userWithPosts, err := getUserPostsWithComments(db, firstUser.ID)
	if err != nil {
		log.Printf("查询用户文章失败: %v", err)
	} else {
		printUserPostsWithComments(userWithPosts)
	}

	// 2. 查询评论数量最多的文章
	mostCommentedPost, err := getMostCommentedPost(db)
	if err != nil {
		log.Printf("查询评论最多文章失败: %v", err)
	} else {
		printMostCommentedPost(mostCommentedPost)
	}
}

func test_CreateComment(db *gorm.DB) {

	comment, err := CreateComment(db, 15, "这是一个测试评论")
	if err != nil {
		log.Fatalf("创建评论失败: %v", err)
	} else {
		log.Printf("创建评论成功: %s (帖子ID: %d)", comment.Content, comment.PostID)
	}

}
func test_createPost(db *gorm.DB) {
	post, err := CreatePost(db, 10, "测试帖子", "这是一个测试帖子")
	if err != nil {
		log.Printf("创建帖子失败: %v", err)
	} else {
		log.Printf("创建帖子成功: %s (用户ID: %d)", post.Title, post.UserID)
	}
	comment, err := CreateComment(db, post.ID, "这是一个测试评论")
	if err != nil {
		log.Fatalf("创建评论失败: %v", err)
	} else {
		log.Printf("创建评论成功: %s (帖子ID: %d)", comment.Content, comment.PostID)
	}
}

// InsertRandomData 生成并插入一条随机测试数据
// 实际业务场景中，数据插入策略取决于具体业务流程
// - 注册新用户：只向User表插入数据
// - 发表文章：向Post表插入数据（需要已存在的UserID）
// - 发表评论：向Comment表插入数据（需要已存在的PostID）
// - 完整流程测试：同时向三张表插入关联数据（如本函数所示）
func InsertRandomData(db *gorm.DB) {
	// 初始化随机数生成器
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	gofakeit.Seed(time.Now().UnixNano())

	// 1. 创建随机用户
	user := User{
		Name:  gofakeit.Name(),
		Email: gofakeit.Email(),
	}
	db.Create(&user)
	log.Printf("创建用户成功: %s (%s)", user.Name, user.Email)

	// 2. 为该用户创建随机帖子
	post := Post{
		Title:   gofakeit.Sentence(r.Intn(2) + 1),
		Content: gofakeit.Paragraph(r.Intn(2)+1, r.Intn(2)+1, r.Intn(2)+10, "\n"),
		UserID:  user.ID,
	}
	db.Create(&post)
	log.Printf("创建帖子成功: %s (用户ID: %d)", post.Title, post.UserID)

	// 3. 为该帖子创建随机评论
	comment := Comment{
		Content: gofakeit.Sentence(r.Intn(10) + 5),
		PostID:  post.ID,
	}
	db.Create(&comment)
	log.Printf("创建评论成功: %s (帖子ID: %d)", comment.Content, comment.PostID)

	log.Println("随机测试数据插入完成")
}

// getUserPostsWithComments 查询某个用户发布的所有文章及其对应的评论信息
func getUserPostsWithComments(db *gorm.DB, userID uint) (User, error) {
	var user User
	err := db.Preload("Posts.Comments").First(&user, userID).Error
	return user, err
}

// getMostCommentedPost 查询评论数量最多的文章信息
func getMostCommentedPost(db *gorm.DB) (Post, error) {
	var post Post
	err := db.Preload("Comments").
		Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON posts.id = comments.post_id").
		Group("posts.id").
		Order("comment_count DESC").
		First(&post).Error
	return post, err
}

// printUserPostsWithComments 打印用户及其文章和评论信息
func printUserPostsWithComments(user User) {
	fmt.Println("\n===== 用户文章及评论信息 =====")
	fmt.Printf("用户ID: %d, 用户名: %s, 邮箱: %s\n", user.ID, user.Name, user.Email)
	fmt.Printf("该用户共有 %d 篇文章\n\n", len(user.Posts))

	for i, post := range user.Posts {
		fmt.Printf("文章 %d:\n", i+1)
		fmt.Printf("  标题: %s\n", post.Title)
		fmt.Printf("  内容: %s\n", post.Content)
		fmt.Printf("  评论数: %d\n", len(post.Comments))

		if len(post.Comments) > 0 {
			fmt.Println("  评论列表:")
			for j, comment := range post.Comments {
				fmt.Printf("    %d. %s\n", j+1, comment.Content)
			}
		}
		fmt.Println()
	}
}

// printMostCommentedPost 打印评论数量最多的文章信息
func printMostCommentedPost(post Post) {
	fmt.Println("\n===== 评论最多的文章 =====")
	fmt.Printf("文章ID: %d\n", post.ID)
	fmt.Printf("标题: %s\n", post.Title)
	fmt.Printf("内容: %s\n", post.Content)
	fmt.Printf("作者ID: %d\n", post.UserID)
	fmt.Printf("评论数量: %d\n\n", len(post.Comments))

	if len(post.Comments) > 0 {
		fmt.Println("评论列表:")
		for i, comment := range post.Comments {
			fmt.Printf("%d. %s\n", i+1, comment.Content)
		}
	}
}

// CreatePost 创建新帖子
func CreatePost(db *gorm.DB, userID uint, title string, content string) (Post, error) {
	// 验证用户是否存在
	var user User
	err := db.First(&user, userID).Error
	if err != nil {
		return Post{}, fmt.Errorf("用户不存在: %w", err)
	}

	// 创建新帖子
	post := Post{
		Title:   title,
		Content: content,
		UserID:  userID,
	}

	// 保存到数据库
	err = db.Create(&post).Error
	if err != nil {
		return Post{}, fmt.Errorf("创建帖子失败: %w", err)
	}

	// 预加载关联的用户信息
	db.Preload("User").First(&post)

	return post, nil
}

// CreateComment 创建新评论
func CreateComment(db *gorm.DB, postID uint, content string) (Comment, error) {
	// 验证帖子是否存在
	var post Post
	err := db.First(&post, postID).Error
	if err != nil {
		return Comment{}, fmt.Errorf("帖子不存在: %w", err)
	}

	// 创建新评论
	comment := Comment{
		Content: content,
		PostID:  postID,
	}

	// 保存到数据库
	err = db.Create(&comment).Error
	if err != nil {
		return Comment{}, fmt.Errorf("创建评论失败: %w", err)
	}

	// 预加载关联的帖子信息
	db.Preload("Post").First(&comment)

	return comment, nil
}
