// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: announcements.sql

package models

import (
	"context"
	"database/sql"
)

const createAnnouncement = `-- name: CreateAnnouncement :one
INSERT INTO
announcement (
    uuid,
    visibility,
    announce_at,
    discord_channel_id,
    discord_message_id
)
VALUES
(?, ?, ?, ?, ?)
RETURNING uuid, visibility, announce_at, discord_channel_id, discord_message_id
`

type CreateAnnouncementParams struct {
	Uuid             string         `json:"uuid"`
	Visibility       string         `json:"visibility"`
	AnnounceAt       int64          `json:"announce_at"`
	DiscordChannelID sql.NullString `json:"discord_channel_id"`
	DiscordMessageID sql.NullString `json:"discord_message_id"`
}

func (q *Queries) CreateAnnouncement(ctx context.Context, arg CreateAnnouncementParams) (Announcement, error) {
	row := q.db.QueryRowContext(ctx, createAnnouncement,
		arg.Uuid,
		arg.Visibility,
		arg.AnnounceAt,
		arg.DiscordChannelID,
		arg.DiscordMessageID,
	)
	var i Announcement
	err := row.Scan(
		&i.Uuid,
		&i.Visibility,
		&i.AnnounceAt,
		&i.DiscordChannelID,
		&i.DiscordMessageID,
	)
	return i, err
}

const getAnnouncement = `-- name: GetAnnouncement :one
SELECT
    uuid,
    visibility,
    announce_at,
    discord_channel_id,
    discord_message_id
FROM
    announcement
WHERE
    uuid = ?
`

func (q *Queries) GetAnnouncement(ctx context.Context, uuid string) (Announcement, error) {
	row := q.db.QueryRowContext(ctx, getAnnouncement, uuid)
	var i Announcement
	err := row.Scan(
		&i.Uuid,
		&i.Visibility,
		&i.AnnounceAt,
		&i.DiscordChannelID,
		&i.DiscordMessageID,
	)
	return i, err
}
