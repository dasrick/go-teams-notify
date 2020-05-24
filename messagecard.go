package goteamsnotify

import (
	"errors"
	"fmt"
	"strings"
)

// MessageCardSectionFact represents a section fact entry that is usually
// displayed in a two-column key/value format.
type MessageCardSectionFact struct {

	// Name is the key for an associated value in a key/value pair
	Name string `json:"name"`

	// Value is the value for an associated key in a key/value pair
	Value string `json:"value"`
}

// MessageCardSectionImage represents an image as used by the heroImage and
// images properties of a section.
type MessageCardSectionImage struct {

	// Image is the URL to the image.
	Image string `json:"image"`

	// Title is a short description of the image. Typically, this description
	// is displayed in a tooltip as the user hovers their mouse over the
	// image.
	Title string `json:"title"`
}

// MessageCardSection represents a section to include in a message card.
type MessageCardSection struct {

	// Title is the title property of a section. This property is displayed
	// in a font that stands out, while not as prominent as the card's title.
	// It is meant to introduce the section and summarize its content,
	// similarly to how the card's title property is meant to summarize the
	// whole card.
	Title string `json:"title,omitempty"`

	// Text is the section's text property. This property is very similar to
	// the text property of the card. It can be used for the same purpose.
	Text string `json:"text,omitempty"`

	// ActivityImage is a property used to display a picture associated with
	// the subject of a message card. For example, this might be the portrait
	// of a person who performed an activity that the message card is
	// associated with.
	ActivityImage string `json:"activityImage,omitempty"`

	// ActivityTitle is a property used to summarize the activity associated
	// with a message card.
	ActivityTitle string `json:"activityTitle,omitempty"`

	// ActivitySubtitle is a property used to show brief, but extended
	// information about an activity associated with a message card. Examples
	// include the date and time the associated activity was taken or the
	// handle of a person associated with the activity.
	ActivitySubtitle string `json:"activitySubtitle,omitempty"`

	// ActivityText is a property used to provide details about the activity.
	// For example, if the message card is used to deliver updates about a
	// topic, then this property would be used to hold the bulk of the content
	// for the update notification.
	ActivityText string `json:"activityText,omitempty"`

	// Markdown represents a toggle to enable or disable Markdown formatting.
	// By default, all text fields in a card and its sections can be formatted
	// using basic Markdown.
	Markdown bool `json:"markdown,omitempty"`

	// StartGroup is the section's startGroup property. This property marks
	// the start of a logical group of information. Typically, sections with
	// startGroup set to true will be visually separated from previous card
	// elements.
	StartGroup bool `json:"startGroup,omitempty"`

	// HeroImage is a property that allows for setting an image as the
	// centerpiece of a message card. This property can also be used to add a
	// banner to the message card.
	// Note: heroImage is not currently supported by Microsoft Teams
	// https://stackoverflow.com/a/45389789
	// We use a pointer to this type in order to have the json package
	// properly omit this field if not explicitly set.
	// https://github.com/golang/go/issues/11939
	// https://stackoverflow.com/questions/18088294/how-to-not-marshal-an-empty-struct-into-json-with-go
	// https://stackoverflow.com/questions/33447334/golang-json-marshal-how-to-omit-empty-nested-struct
	HeroImage *MessageCardSectionImage `json:"heroImage,omitempty"`

	// Facts is a collection of MessageCardSectionFact values. A section entry
	// usually is displayed in a two-column key/value format.
	Facts []MessageCardSectionFact `json:"facts,omitempty"`

	// Images is a property that allows for the inclusion of a photo gallery
	// inside a section.
	// We use a slice of pointers to this type in order to have the json
	// package properly omit this field if not explicitly set.
	// https://github.com/golang/go/issues/11939
	// https://stackoverflow.com/questions/18088294/how-to-not-marshal-an-empty-struct-into-json-with-go
	// https://stackoverflow.com/questions/33447334/golang-json-marshal-how-to-omit-empty-nested-struct
	Images []*MessageCardSectionImage `json:"images,omitempty"`
}

// MessageCard represents a legacy actionable message card used via Office 365
// or Microsoft Teams connectors.
type MessageCard struct {
	// Required; must be set to "MessageCard"
	Type string `json:"@type"`

	// Required; must be set to "https://schema.org/extensions"
	Context string `json:"@context"`

	// Summary is required if the card does not contain a text property,
	// otherwise optional. The summary property is typically displayed in the
	// list view in Outlook, as a way to quickly determine what the card is
	// all about. Summary appears to only be used when there are sections defined
	Summary string `json:"summary,omitempty"`

	// Title is the title property of a card. is meant to be rendered in a
	// prominent way, at the very top of the card. Use it to introduce the
	// content of the card in such a way users will immediately know what to
	// expect.
	Title string `json:"title,omitempty"`

	// Text is required if the card does not contain a summary property,
	// otherwise optional. The text property is meant to be displayed in a
	// normal font below the card's title. Use it to display content, such as
	// the description of the entity being referenced, or an abstract of a
	// news article.
	Text string `json:"text,omitempty"`

	// Specifies a custom brand color for the card. The color will be
	// displayed in a non-obtrusive manner.
	ThemeColor string `json:"themeColor,omitempty"`

	// Sections is a collection of sections to include in the card.
	Sections []*MessageCardSection `json:"sections,omitempty"`
}

// AddSection adds one or many additional MessageCardSection values to a
// MessageCard. Validation is performed to reject invalid values with an error
// message.
func (mc *MessageCard) AddSection(section ...*MessageCardSection) error {
	for _, s := range section {
		// bail if a completely nil section provided
		if s == nil {
			return fmt.Errorf("func AddSection: nil MessageCardSection received")
		}

		// Perform validation of all MessageCardSection fields in an effort to
		// avoid adding a MessageCardSection with zero value fields. This is
		// done to avoid generating an empty sections JSON array since the
		// Sections slice for the MessageCard type would technically not be at
		// a zero value state. Due to this non-zero value state, the
		// encoding/json package would end up including the Sections struct
		// field in the output JSON.
		// See also https://github.com/golang/go/issues/11939
		switch {
		// If any of these cases trigger, skip over the `default` case
		// statement and add the section.
		case s.Images != nil:
		case s.Facts != nil:
		case s.HeroImage != nil:
		case s.StartGroup:
		case s.Markdown:
		case s.ActivityText != "":
		case s.ActivitySubtitle != "":
		case s.ActivityTitle != "":
		case s.ActivityImage != "":
		case s.Text != "":
		case s.Title != "":

		default:
			return fmt.Errorf("all fields found to be at zero-value, skipping section")
		}

		mc.Sections = append(mc.Sections, s)
	}

	return nil
}

// AddFact adds one or many additional MessageCardSectionFact values to a
// MessageCardSection
func (mcs *MessageCardSection) AddFact(fact ...MessageCardSectionFact) error {
	for _, f := range fact {
		if f.Name == "" {
			return fmt.Errorf("empty Name field received for new fact: %+v", f)
		}

		if f.Value == "" {
			return fmt.Errorf("empty Name field received for new fact: %+v", f)
		}
	}

	mcs.Facts = append(mcs.Facts, fact...)

	return nil
}

// AddFactFromKeyValue accepts a key and slice of values and converts them to
// MessageCardSectionFact values
func (mcs *MessageCardSection) AddFactFromKeyValue(key string, values ...string) error {
	// validate arguments

	if key == "" {
		return errors.New("empty key received for new fact")
	}

	if len(values) < 1 {
		return errors.New("no values received for new fact")
	}

	fact := MessageCardSectionFact{
		Name:  key,
		Value: strings.Join(values, ", "),
	}
	// TODO: Explicitly define or use constructor?
	// fact := NewMessageCardSectionFact()
	// fact.Name = key
	// fact.Value = strings.Join(values, ", ")

	mcs.Facts = append(mcs.Facts, fact)

	// if we made it this far then all should be well
	return nil
}

// AddImage adds an image to a MessageCard section. These images are used to
// provide a photo gallery inside a MessageCard section.
func (mcs *MessageCardSection) AddImage(sectionImage ...MessageCardSectionImage) error {
	for i := range sectionImage {
		if sectionImage[i].Image == "" {
			return fmt.Errorf("cannot add empty image URL")
		}

		if sectionImage[i].Title == "" {
			return fmt.Errorf("cannot add empty image title")
		}

		mcs.Images = append(mcs.Images, &sectionImage[i])
	}

	return nil
}

// AddHeroImageStr adds a Hero Image to a MessageCard section using string
// arguments. This image is used as the centerpiece or banner of a message
// card.
func (mcs *MessageCardSection) AddHeroImageStr(imageURL string, imageTitle string) error {
	if imageURL == "" {
		return fmt.Errorf("cannot add empty hero image URL")
	}

	if imageTitle == "" {
		return fmt.Errorf("cannot add empty hero image title")
	}

	heroImage := MessageCardSectionImage{
		Image: imageURL,
		Title: imageTitle,
	}
	// TODO: Explicitly define or use constructor?
	// heroImage := NewMessageCardSectionImage()
	// heroImage.Image = imageURL
	// heroImage.Title = imageTitle

	mcs.HeroImage = &heroImage

	// our validation checks didn't find any problems
	return nil
}

// AddHeroImage adds a Hero Image to a MessageCard section using a
// MessageCardSectionImage argument. This image is used as the centerpiece or
// banner of a message card.
func (mcs *MessageCardSection) AddHeroImage(heroImage MessageCardSectionImage) error {
	if heroImage.Image == "" {
		return fmt.Errorf("cannot add empty hero image URL")
	}

	if heroImage.Title == "" {
		return fmt.Errorf("cannot add empty hero image title")
	}

	mcs.HeroImage = &heroImage

	// our validation checks didn't find any problems
	return nil
}

// NewMessageCard creates a new message card with fields required by the
// legacy message card format already predefined
func NewMessageCard() MessageCard {
	// define expected values to meet Office 365 Connector card requirements
	// https://docs.microsoft.com/en-us/outlook/actionable-messages/message-card-reference#card-fields
	msgCard := MessageCard{
		Type:    "MessageCard",
		Context: "https://schema.org/extensions",
	}

	return msgCard
}

// NewMessageCardSection creates an empty message card section
func NewMessageCardSection() *MessageCardSection {
	msgCardSection := MessageCardSection{}
	return &msgCardSection
}

// NewMessageCardSectionFact creates an empty message card section fact
func NewMessageCardSectionFact() MessageCardSectionFact {
	msgCardSectionFact := MessageCardSectionFact{}
	return msgCardSectionFact
}

// NewMessageCardSectionImage creates an empty image for use with message card
// section
func NewMessageCardSectionImage() MessageCardSectionImage {
	msgCardSectionImage := MessageCardSectionImage{}
	return msgCardSectionImage
}
