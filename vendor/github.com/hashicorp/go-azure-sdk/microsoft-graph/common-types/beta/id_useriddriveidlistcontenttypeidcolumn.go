package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdListContentTypeIdColumnId{}

// UserIdDriveIdListContentTypeIdColumnId is a struct representing the Resource ID for a User Id Drive Id List Content Type Id Column
type UserIdDriveIdListContentTypeIdColumnId struct {
	UserId             string
	DriveId            string
	ContentTypeId      string
	ColumnDefinitionId string
}

// NewUserIdDriveIdListContentTypeIdColumnID returns a new UserIdDriveIdListContentTypeIdColumnId struct
func NewUserIdDriveIdListContentTypeIdColumnID(userId string, driveId string, contentTypeId string, columnDefinitionId string) UserIdDriveIdListContentTypeIdColumnId {
	return UserIdDriveIdListContentTypeIdColumnId{
		UserId:             userId,
		DriveId:            driveId,
		ContentTypeId:      contentTypeId,
		ColumnDefinitionId: columnDefinitionId,
	}
}

// ParseUserIdDriveIdListContentTypeIdColumnID parses 'input' into a UserIdDriveIdListContentTypeIdColumnId
func ParseUserIdDriveIdListContentTypeIdColumnID(input string) (*UserIdDriveIdListContentTypeIdColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListContentTypeIdColumnId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListContentTypeIdColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdListContentTypeIdColumnIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdListContentTypeIdColumnId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdListContentTypeIdColumnIDInsensitively(input string) (*UserIdDriveIdListContentTypeIdColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListContentTypeIdColumnId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListContentTypeIdColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdListContentTypeIdColumnId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ContentTypeId, ok = input.Parsed["contentTypeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId", input)
	}

	if id.ColumnDefinitionId, ok = input.Parsed["columnDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "columnDefinitionId", input)
	}

	return nil
}

// ValidateUserIdDriveIdListContentTypeIdColumnID checks that 'input' can be parsed as a User Id Drive Id List Content Type Id Column ID
func ValidateUserIdDriveIdListContentTypeIdColumnID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdListContentTypeIdColumnID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id List Content Type Id Column ID
func (id UserIdDriveIdListContentTypeIdColumnId) ID() string {
	fmtString := "/users/%s/drives/%s/list/contentTypes/%s/columns/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.ContentTypeId, id.ColumnDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id List Content Type Id Column ID
func (id UserIdDriveIdListContentTypeIdColumnId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
		resourceids.StaticSegment("columns", "columns", "columns"),
		resourceids.UserSpecifiedSegment("columnDefinitionId", "columnDefinitionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id List Content Type Id Column ID
func (id UserIdDriveIdListContentTypeIdColumnId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Column Definition: %q", id.ColumnDefinitionId),
	}
	return fmt.Sprintf("User Id Drive Id List Content Type Id Column (%s)", strings.Join(components, "\n"))
}