# cmsblocks

[![Tests Status](https://github.com/dracory/cmsblocks/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/dracory/cmsblocks/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dracory/cmsblocks)](https://goreportcard.com/report/github.com/dracory/cmsblocks)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/dracory/cmsblocks)](https://pkg.go.dev/github.com/dracory/cmsblocks)

Reusable `dracory/cmsstore` custom block types for Blueprint-style projects.

## Block types

| Package | Type key | Renders |
| --- | --- | --- |
| `blogpost` | `blog_post` | A single blog post (auto-detected from URL or fixed `post_id`), with tags, share links, prev/next navigation |
| `blogpostlist` | `blog_post_list` | Paginated blog post cards, column layout, excerpts |
| `search` | `search` | Search box + paginated results across CMS pages and blog posts |

## Usage

```go
postURL := func(postID, postSlug string) string {
    return links.Website().BlogPost(postID, postSlug) // or any project URL builder
}

b := blogpost.NewBlogPostBlockType(blogStore)
b.SetPostURLFunc(postURL) // optional — defaults to /blog/post/{id}/{slug}
cmsstore.RegisterCustomBlockType(b)
```

Each block type accepts an optional `SetPostURLFunc` so the project controls
the public post URL format; without it, links default to
`/blog/post/{id}/{slug}`.

## Dependencies

`dracory/blogstore`, `dracory/cmsstore`, `dracory/form`, `dracory/hb`,
`dracory/social`, `samber/lo`, `spf13/cast`.
