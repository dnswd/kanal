# Kanal
## Overview

## Endpoints
### Topic
- [ ] `GET    /news` List every news
- [ ] `POST   /news` Add new topic
- [ ] `GET    /news/:topic` List every news by topic
- [ ] `PUT    /news/:topic` Modify a topic name
- [ ] Optional query param `status=draft,deleted,published`

### News
- [ ] `POST   /article` Submit a news article draft
- [ ] `GET    /article/:id` Get the article of the news
- [ ] `PUT    /article/:id` Replace a news article
- [ ] `DELETE /article/:id` Delete a news article
  - [ ] Query param `hard=true,false`, default to `false`
- [ ] `GET    /article/:id/tags` Get tags of a news article
- [ ] `PUT    /article/:id/tags` Replace tag of existing article
- [ ] `GET    /article/:id/author` Get author of a news article
- [ ] `POST   /article/:id/publish` Publish a news article <- unsure
- [ ] `POST   /article/:id/unpublish` Set status of a news article to `draft`

### Tags
- [ ] `GET    /tags` Get all available tags
- [ ] `POST   /tags` Create a new tag
- [ ] `PUT    /tags/:id` Rename a tag
- [ ] `DELETE /tags/:id` Delete a tag

## Models
### Topic
- Several news belongs to a topic
- Has an id and a name

### News
- Has an id, title, author, editor, published date, and article body
- Each news belongs only to one topic
- Each news can have one or more Tags

### Tags
- Has an id and a name
- Many tags belongs to many news

### ER Diagram
[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgVE9QSUMgfHwuLm97IE5FV1MgOiBcImJlbG9uZ3MgdG9cIlxuICAgIE5FV1MgfXwgLi4gb3sgVEFHUyA6IFwiaGFzXCJcblxuICAgICAgICAgICAgIiwibWVybWFpZCI6e30sInVwZGF0ZUVkaXRvciI6ZmFsc2V9)](https://mermaid-js.github.io/mermaid-live-editor/#/edit/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgVE9QSUMgfHwuLm97IE5FV1MgOiBcImJlbG9uZ3MgdG9cIlxuICAgIE5FV1MgfXwgLi4gb3sgVEFHUyA6IFwiaGFzXCJcblxuICAgICAgICAgICAgIiwibWVybWFpZCI6e30sInVwZGF0ZUVkaXRvciI6ZmFsc2V9)

```
erDiagram
    TOPIC ||..o{ NEWS : "belongs to"
    NEWS }| .. o{ TAGS : "has"
```

## Assumptions
- Assuming each news belongs to only one topic
- Assuming we're building for a news company, so we don't have to record `publisher` information on News model.
- Assuming we don't have to worry about multimedia objects/blobs (e.g. videos, images, etc)
- Assuming we don't have to worry about links on the News's article

## TODO
- [X] DB Models
  - [X] Topic
  - [X] News
  - [X] Tags
- [ ] Endpoints
- [ ] Tests