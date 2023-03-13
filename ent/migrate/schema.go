// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArticlesColumns holds the columns for the "articles" table.
	ArticlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString, Size: 1024},
		{Name: "slug", Type: field.TypeString, Unique: true, Size: 1024},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 512},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
	}
	// ArticlesTable holds the schema information for the "articles" table.
	ArticlesTable = &schema.Table{
		Name:       "articles",
		Columns:    ArticlesColumns,
		PrimaryKey: []*schema.Column{ArticlesColumns[0]},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 216},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "content", Type: field.TypeString, Size: 512},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
	}
	// GalleriesColumns holds the columns for the "galleries" table.
	GalleriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 216},
	}
	// GalleriesTable holds the schema information for the "galleries" table.
	GalleriesTable = &schema.Table{
		Name:       "galleries",
		Columns:    GalleriesColumns,
		PrimaryKey: []*schema.Column{GalleriesColumns[0]},
	}
	// MetadataColumns holds the columns for the "metadata" table.
	MetadataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Size: 216},
	}
	// MetadataTable holds the schema information for the "metadata" table.
	MetadataTable = &schema.Table{
		Name:       "metadata",
		Columns:    MetadataColumns,
		PrimaryKey: []*schema.Column{MetadataColumns[0]},
	}
	// NewslettersColumns holds the columns for the "newsletters" table.
	NewslettersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "message", Type: field.TypeString, Size: 1024},
	}
	// NewslettersTable holds the schema information for the "newsletters" table.
	NewslettersTable = &schema.Table{
		Name:       "newsletters",
		Columns:    NewslettersColumns,
		PrimaryKey: []*schema.Column{NewslettersColumns[0]},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 216},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 64},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString, Size: 264},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"admin", "author", "user", "subscriber"}, Default: "user"},
		{Name: "token", Type: field.TypeString, Nullable: true},
		{Name: "token_expired", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// ArticleCategoriesColumns holds the columns for the "article_categories" table.
	ArticleCategoriesColumns = []*schema.Column{
		{Name: "article_id", Type: field.TypeUUID},
		{Name: "category_id", Type: field.TypeUUID},
	}
	// ArticleCategoriesTable holds the schema information for the "article_categories" table.
	ArticleCategoriesTable = &schema.Table{
		Name:       "article_categories",
		Columns:    ArticleCategoriesColumns,
		PrimaryKey: []*schema.Column{ArticleCategoriesColumns[0], ArticleCategoriesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "article_categories_article_id",
				Columns:    []*schema.Column{ArticleCategoriesColumns[0]},
				RefColumns: []*schema.Column{ArticlesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "article_categories_category_id",
				Columns:    []*schema.Column{ArticleCategoriesColumns[1]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArticlesTable,
		CategoriesTable,
		CommentsTable,
		GalleriesTable,
		MetadataTable,
		NewslettersTable,
		TagsTable,
		UsersTable,
		ArticleCategoriesTable,
	}
)

func init() {
	ArticleCategoriesTable.ForeignKeys[0].RefTable = ArticlesTable
	ArticleCategoriesTable.ForeignKeys[1].RefTable = CategoriesTable
}
