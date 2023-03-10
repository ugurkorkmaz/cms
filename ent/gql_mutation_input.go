// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/user"
	"time"

	"github.com/google/uuid"
)

// CreateArticleInput represents a mutation input for creating articles.
type CreateArticleInput struct {
	Title       string
	Slug        string
	Description *string
	Content     string
	CategoryIDs []uuid.UUID
}

// Mutate applies the CreateArticleInput on the ArticleMutation builder.
func (i *CreateArticleInput) Mutate(m *ArticleMutation) {
	m.SetTitle(i.Title)
	m.SetSlug(i.Slug)
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	m.SetContent(i.Content)
	if v := i.CategoryIDs; len(v) > 0 {
		m.AddCategoryIDs(v...)
	}
}

// SetInput applies the change-set in the CreateArticleInput on the ArticleCreate builder.
func (c *ArticleCreate) SetInput(i CreateArticleInput) *ArticleCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateArticleInput represents a mutation input for updating articles.
type UpdateArticleInput struct {
	Title             *string
	Slug              *string
	ClearDescription  bool
	Description       *string
	Content           *string
	ClearCategories   bool
	AddCategoryIDs    []uuid.UUID
	RemoveCategoryIDs []uuid.UUID
}

// Mutate applies the UpdateArticleInput on the ArticleMutation builder.
func (i *UpdateArticleInput) Mutate(m *ArticleMutation) {
	if v := i.Title; v != nil {
		m.SetTitle(*v)
	}
	if v := i.Slug; v != nil {
		m.SetSlug(*v)
	}
	if i.ClearDescription {
		m.ClearDescription()
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if v := i.Content; v != nil {
		m.SetContent(*v)
	}
	if i.ClearCategories {
		m.ClearCategories()
	}
	if v := i.AddCategoryIDs; len(v) > 0 {
		m.AddCategoryIDs(v...)
	}
	if v := i.RemoveCategoryIDs; len(v) > 0 {
		m.RemoveCategoryIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateArticleInput on the ArticleUpdate builder.
func (c *ArticleUpdate) SetInput(i UpdateArticleInput) *ArticleUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateArticleInput on the ArticleUpdateOne builder.
func (c *ArticleUpdateOne) SetInput(i UpdateArticleInput) *ArticleUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateCategoryInput represents a mutation input for creating categories.
type CreateCategoryInput struct {
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	Name       string
	ArticleIDs []uuid.UUID
}

// Mutate applies the CreateCategoryInput on the CategoryMutation builder.
func (i *CreateCategoryInput) Mutate(m *CategoryMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetName(i.Name)
	if v := i.ArticleIDs; len(v) > 0 {
		m.AddArticleIDs(v...)
	}
}

// SetInput applies the change-set in the CreateCategoryInput on the CategoryCreate builder.
func (c *CategoryCreate) SetInput(i CreateCategoryInput) *CategoryCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateCategoryInput represents a mutation input for updating categories.
type UpdateCategoryInput struct {
	UpdatedAt        *time.Time
	Name             *string
	ClearArticle     bool
	AddArticleIDs    []uuid.UUID
	RemoveArticleIDs []uuid.UUID
}

// Mutate applies the UpdateCategoryInput on the CategoryMutation builder.
func (i *UpdateCategoryInput) Mutate(m *CategoryMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearArticle {
		m.ClearArticle()
	}
	if v := i.AddArticleIDs; len(v) > 0 {
		m.AddArticleIDs(v...)
	}
	if v := i.RemoveArticleIDs; len(v) > 0 {
		m.RemoveArticleIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateCategoryInput on the CategoryUpdate builder.
func (c *CategoryUpdate) SetInput(i UpdateCategoryInput) *CategoryUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateCategoryInput on the CategoryUpdateOne builder.
func (c *CategoryUpdateOne) SetInput(i UpdateCategoryInput) *CategoryUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateCommentInput represents a mutation input for creating comments.
type CreateCommentInput struct {
	CreatedAt *time.Time
	Content   string
}

// Mutate applies the CreateCommentInput on the CommentMutation builder.
func (i *CreateCommentInput) Mutate(m *CommentMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	m.SetContent(i.Content)
}

// SetInput applies the change-set in the CreateCommentInput on the CommentCreate builder.
func (c *CommentCreate) SetInput(i CreateCommentInput) *CommentCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateCommentInput represents a mutation input for updating comments.
type UpdateCommentInput struct {
	Content *string
}

// Mutate applies the UpdateCommentInput on the CommentMutation builder.
func (i *UpdateCommentInput) Mutate(m *CommentMutation) {
	if v := i.Content; v != nil {
		m.SetContent(*v)
	}
}

// SetInput applies the change-set in the UpdateCommentInput on the CommentUpdate builder.
func (c *CommentUpdate) SetInput(i UpdateCommentInput) *CommentUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateCommentInput on the CommentUpdateOne builder.
func (c *CommentUpdateOne) SetInput(i UpdateCommentInput) *CommentUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateGalleryInput represents a mutation input for creating galleries.
type CreateGalleryInput struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Name      string
}

// Mutate applies the CreateGalleryInput on the GalleryMutation builder.
func (i *CreateGalleryInput) Mutate(m *GalleryMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetName(i.Name)
}

// SetInput applies the change-set in the CreateGalleryInput on the GalleryCreate builder.
func (c *GalleryCreate) SetInput(i CreateGalleryInput) *GalleryCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateGalleryInput represents a mutation input for updating galleries.
type UpdateGalleryInput struct {
	UpdatedAt *time.Time
	Name      *string
}

// Mutate applies the UpdateGalleryInput on the GalleryMutation builder.
func (i *UpdateGalleryInput) Mutate(m *GalleryMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
}

// SetInput applies the change-set in the UpdateGalleryInput on the GalleryUpdate builder.
func (c *GalleryUpdate) SetInput(i UpdateGalleryInput) *GalleryUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateGalleryInput on the GalleryUpdateOne builder.
func (c *GalleryUpdateOne) SetInput(i UpdateGalleryInput) *GalleryUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateMetadataInput represents a mutation input for creating metadataslice.
type CreateMetadataInput struct {
	UpdatedAt *time.Time
	CreatedAt *time.Time
	Title     string
}

// Mutate applies the CreateMetadataInput on the MetadataMutation builder.
func (i *CreateMetadataInput) Mutate(m *MetadataMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	m.SetTitle(i.Title)
}

// SetInput applies the change-set in the CreateMetadataInput on the MetadataCreate builder.
func (c *MetadataCreate) SetInput(i CreateMetadataInput) *MetadataCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateMetadataInput represents a mutation input for updating metadataslice.
type UpdateMetadataInput struct {
	UpdatedAt *time.Time
	Title     *string
}

// Mutate applies the UpdateMetadataInput on the MetadataMutation builder.
func (i *UpdateMetadataInput) Mutate(m *MetadataMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Title; v != nil {
		m.SetTitle(*v)
	}
}

// SetInput applies the change-set in the UpdateMetadataInput on the MetadataUpdate builder.
func (c *MetadataUpdate) SetInput(i UpdateMetadataInput) *MetadataUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateMetadataInput on the MetadataUpdateOne builder.
func (c *MetadataUpdateOne) SetInput(i UpdateMetadataInput) *MetadataUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateNewsletterInput represents a mutation input for creating newsletters.
type CreateNewsletterInput struct {
	UpdatedAt *time.Time
	CreatedAt *time.Time
	Message   string
}

// Mutate applies the CreateNewsletterInput on the NewsletterMutation builder.
func (i *CreateNewsletterInput) Mutate(m *NewsletterMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	m.SetMessage(i.Message)
}

// SetInput applies the change-set in the CreateNewsletterInput on the NewsletterCreate builder.
func (c *NewsletterCreate) SetInput(i CreateNewsletterInput) *NewsletterCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateNewsletterInput represents a mutation input for updating newsletters.
type UpdateNewsletterInput struct {
	UpdatedAt *time.Time
	Message   *string
}

// Mutate applies the UpdateNewsletterInput on the NewsletterMutation builder.
func (i *UpdateNewsletterInput) Mutate(m *NewsletterMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Message; v != nil {
		m.SetMessage(*v)
	}
}

// SetInput applies the change-set in the UpdateNewsletterInput on the NewsletterUpdate builder.
func (c *NewsletterUpdate) SetInput(i UpdateNewsletterInput) *NewsletterUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateNewsletterInput on the NewsletterUpdateOne builder.
func (c *NewsletterUpdateOne) SetInput(i UpdateNewsletterInput) *NewsletterUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateTagInput represents a mutation input for creating tags.
type CreateTagInput struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Name      string
}

// Mutate applies the CreateTagInput on the TagMutation builder.
func (i *CreateTagInput) Mutate(m *TagMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetName(i.Name)
}

// SetInput applies the change-set in the CreateTagInput on the TagCreate builder.
func (c *TagCreate) SetInput(i CreateTagInput) *TagCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateTagInput represents a mutation input for updating tags.
type UpdateTagInput struct {
	UpdatedAt *time.Time
	Name      *string
}

// Mutate applies the UpdateTagInput on the TagMutation builder.
func (i *UpdateTagInput) Mutate(m *TagMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
}

// SetInput applies the change-set in the UpdateTagInput on the TagUpdate builder.
func (c *TagUpdate) SetInput(i UpdateTagInput) *TagUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateTagInput on the TagUpdateOne builder.
func (c *TagUpdateOne) SetInput(i UpdateTagInput) *TagUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	UpdatedAt    *time.Time
	CreatedAt    *time.Time
	Name         string
	Email        string
	Password     string
	Role         *user.Role
	Token        *string
	TokenExpired *time.Time
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	m.SetName(i.Name)
	m.SetEmail(i.Email)
	m.SetPassword(i.Password)
	if v := i.Role; v != nil {
		m.SetRole(*v)
	}
	if v := i.Token; v != nil {
		m.SetToken(*v)
	}
	if v := i.TokenExpired; v != nil {
		m.SetTokenExpired(*v)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	UpdatedAt         *time.Time
	Name              *string
	Password          *string
	Role              *user.Role
	ClearToken        bool
	Token             *string
	ClearTokenExpired bool
	TokenExpired      *time.Time
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Password; v != nil {
		m.SetPassword(*v)
	}
	if v := i.Role; v != nil {
		m.SetRole(*v)
	}
	if i.ClearToken {
		m.ClearToken()
	}
	if v := i.Token; v != nil {
		m.SetToken(*v)
	}
	if i.ClearTokenExpired {
		m.ClearTokenExpired()
	}
	if v := i.TokenExpired; v != nil {
		m.SetTokenExpired(*v)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
