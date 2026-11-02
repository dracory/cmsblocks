// Package teststore builds real in-memory SQLite-backed stores for tests.
package teststore

import (
	"database/sql"
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/dracory/blogstore"
	"github.com/dracory/cmsstore"
	_ "modernc.org/sqlite"
)

var counter atomic.Int64

// newDB opens a unique in-memory SQLite database for the test.
func newDB(t *testing.T) *sql.DB {
	t.Helper()
	dsn := fmt.Sprintf("file:cmsblocks_test_%d?mode=memory&cache=shared", counter.Add(1))
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		t.Fatalf("failed to open sqlite: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })
	return db
}

// NewBlogStore returns a migrated blog store backed by in-memory SQLite.
func NewBlogStore(t *testing.T) blogstore.StoreInterface {
	t.Helper()
	st, err := blogstore.NewStore(blogstore.NewStoreOptions{
		DB:                  newDB(t),
		PostTableName:       "test_blog_post",
		TaxonomyEnabled:     true,
		TaxonomyTableName:   "test_blog_taxonomy",
		TermTableName:       "test_blog_term",
		VersioningEnabled:   true,
		VersioningTableName: "test_blog_version",
		AutomigrateEnabled:  true,
	})
	if err != nil {
		t.Fatalf("failed to create blog store: %v", err)
	}
	return st
}

// NewCmsStore returns a CMS store backed by in-memory SQLite.
func NewCmsStore(t *testing.T) cmsstore.StoreInterface {
	t.Helper()
	st, err := cmsstore.NewStore(cmsstore.NewStoreOptions{
		DB:                   newDB(t),
		BlockTableName:       "test_cms_block",
		PageTableName:        "test_cms_page",
		TemplateTableName:    "test_cms_template",
		SiteTableName:        "test_cms_site",
		MenusEnabled:         true,
		MenuItemTableName:    "test_cms_menu_item",
		MenuTableName:        "test_cms_menu",
		TranslationsEnabled:  true,
		TranslationTableName: "test_cms_translation",
		VersioningEnabled:    true,
		VersioningTableName:  "test_cms_version",
	})
	if err != nil {
		t.Fatalf("failed to create cms store: %v", err)
	}
	if err := st.AutoMigrate(t.Context()); err != nil {
		t.Fatalf("failed to migrate cms store: %v", err)
	}
	return st
}
