package models

import "github.com/jinzhu/gorm"

//日志数据结构
type Blog struct {
	Gid         int64  `json:"gid",gorm:"primary_key;auto_increment"`
	Title       string `json:"title"`
	Date        int64  `json:"-"`
	Content     string `json:"content"`
	Excerpt     string `json:"excerpt"`
	Alias       string `json:"alias"`
	Author      int64  `json:"author"`
	Sortid      int64  `json:"sortid"`
	Type        string `json:"type"`
	Views       int    `json:"views"`
	Comnum      int    `json:"comnun"`
	Attnum      int    `json:"attnum"`
	Top         string `json:"top"`
	Sortop      string `json:"sortop"`
	Hide        string `json:"hide"`
	Checked     string `json:"checked"`
	AllowRemark string `json:"all_remark"` //是否允许评论
	Password    string `json:"-"`
	Template    string `json:"template"`

	DateFmt string `json:"date_fmt"`
	User    User   `gorm:"foreignkey:Author;association_foreignkey:Uid"`
	Sort    Sort   `gorm:"foreignkey:Sortid;association_foreignkey:Sid"`
}

//归档数据结构,显示效果 2019年4月(2)
type Record struct {
	Ym    string `json:"ym"`
	Count string `json:"count"`
}

//指定表名
//func (m *Blog) TableName() string {
//	return "e_blog"
//}

//通过ID获取一篇文章
func GetBlogById(gid int64) (Blog, error) {
	var blog Blog
	err := Orm.Where(&Blog{Gid: gid}).Preload("User").Preload("Sort").First(&blog).Error
	if err == nil {
		blog.Views++
		Orm.Model(&blog).UpdateColumn("views", gorm.Expr("views + ?", 1))
	}
	return blog, err
}

//获取所有文章
func GetBlogList(page int64, keyword string, sortid int64) []Blog {
	if page < 1 {
		page = 1
	}
	var blogList []Blog
	//var count int64
	var offset int64
	var limit int64 = 10
	offset = (page - 1) * limit
	where := &Blog{}
	where.Hide = "n"
	where.Checked = "y"
	if sortid > 0 {
		//查询sort的id
		where.Sortid = sortid
	}
	qs := Orm.Where(where).Preload("User").Preload("Sort").Limit(limit).Offset(offset).Order("top, date desc")
	if len(keyword) > 0 {
		qs = qs.Where("title LIKE ?", "%"+keyword+"%")
	}
	qs.Find(&blogList)
	return blogList
}

//日志归档数据
func GetRecord() []Record {
	sql := "select from_unixtime(`date`, '%Y年%m月') as ym, count(*) as count from e_blog where hide = 'n' and checked = 'y' and type = 'blog' group by ym order by `ym` desc"
	var records []Record
	Orm.Raw(sql).Find(&records)
	return records
}
