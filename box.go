package fyne_widget

import (
	"math"

	"fyne.io/fyne/v2"
)

const (
	SizeInfinity float32 = math.MaxFloat32
)

type BoxConstraints struct {
	// The minimum width that satisfies the constraints.
	MinWidth float32

	// The maximum width that satisfies the constraints.
	// Might be [double.infinity].
	MaxWidth float32

	// The minimum height that satisfies the constraints.
	MinHeight float32

	// The maximum height that satisfies the constraints.
	// Might be [double.infinity].
	MaxHeight float32
}

func NewBoxConstraints(minWidth, maxWidth, minHeight, maxHeight float32) BoxConstraints {
	return BoxConstraints{
		MinWidth:  minWidth,
		MaxWidth:  maxWidth,
		MinHeight: minHeight,
		MaxHeight: maxHeight,
	}
}

func DefaultBoxConstraints() BoxConstraints {
	return BoxConstraints{
		MinWidth:  0,
		MaxWidth:  SizeInfinity,
		MinHeight: 0,
		MaxHeight: SizeInfinity,
	}
}

// Creates box constraints that is respected only by the given size.
func TightBoxConstraints(size fyne.Size) BoxConstraints {
	return BoxConstraints{
		MinWidth:  size.Width,
		MaxWidth:  size.Width,
		MinHeight: size.Height,
		MaxHeight: size.Height,
	}
}

// Return box constraints that is width respected only by the given size.
func (o BoxConstraints) TightWidth(width float32) BoxConstraints {
	r := o
	r.MinWidth = width
	r.MaxWidth = width
	return r
}

// Return box constraints that is height respected only by the given size.
func (o BoxConstraints) TightHeight(height float32) BoxConstraints {
	r := o
	r.MinHeight = height
	r.MaxHeight = height
	return r
}

// Creates box constraints that require the given width or height, except if
// they are infinite.
func TightForFiniteBoxContraints(width float32, height float32) BoxConstraints {
	box := BoxConstraints{
		MinWidth:  width,
		MaxWidth:  width,
		MinHeight: height,
		MaxHeight: height,
	}
	if width == SizeInfinity {
		box.MinWidth = 0
	}
	if height == SizeInfinity {
		box.MinHeight = 0
	}
	return box
}

// Creates box constraints that forbid sizes larger than the given size.
func LooseBoxConstraints(size fyne.Size) BoxConstraints {
	return BoxConstraints{
		MinWidth:  0,
		MaxWidth:  size.Width,
		MinHeight: 0,
		MaxHeight: size.Height,
	}
}

/// Returns new box constraints that are smaller by the given edge dimensions.
// BoxConstraints deflate(EdgeInsets edges) {
// 	assert(debugAssertIsValid());
// 	final double horizontal = edges.horizontal;
// 	final double vertical = edges.vertical;
// 	final double deflatedMinWidth = math.max(0.0, minWidth - horizontal);
// 	final double deflatedMinHeight = math.max(0.0, minHeight - vertical);
// 	return BoxConstraints(
// 	  minWidth: deflatedMinWidth,
// 	  maxWidth: math.max(deflatedMinWidth, maxWidth - horizontal),
// 	  minHeight: deflatedMinHeight,
// 	  maxHeight: math.max(deflatedMinHeight, maxHeight - vertical),
// 	);
//   }

// Returns new box constraints that remove the minimum width and height requirements.
func (o BoxConstraints) loosen() BoxConstraints {
	r := o
	r.MinWidth = 0
	r.MinHeight = 0
	return r
}

func ClampFloat32(x float32, min float32, max float32) float32 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	if x == SizeInfinity {
		return max
	}
	return x
}

// Returns new box constraints that respect the given constraints while being
// as close as possible to the original constraints.
func (o BoxConstraints) enforce(constraints BoxConstraints) BoxConstraints {
	return BoxConstraints{
		MinWidth:  ClampFloat32(o.MinWidth, constraints.MinWidth, constraints.MaxWidth),
		MaxWidth:  ClampFloat32(o.MaxWidth, constraints.MinWidth, constraints.MaxWidth),
		MinHeight: ClampFloat32(o.MinHeight, constraints.MinHeight, constraints.MaxHeight),
		MaxHeight: ClampFloat32(o.MaxHeight, constraints.MinHeight, constraints.MaxHeight),
	}
}

// Returns new box constraints with a tight width as close to
// the given width and height as possible while still respecting the original
// box constraints.
func (o BoxConstraints) TightenWidth(v float32) BoxConstraints {
	return BoxConstraints{
		MinWidth:  ClampFloat32(v, o.MinWidth, o.MaxWidth),
		MaxWidth:  ClampFloat32(v, o.MinWidth, o.MaxWidth),
		MinHeight: o.MinHeight,
		MaxHeight: o.MaxHeight,
	}
}

// Returns new box constraints with a tight height as close to
// the given width and height as possible while still respecting the original
// box constraints.
func (o BoxConstraints) TightenHeight(v float32) BoxConstraints {
	return BoxConstraints{
		MinWidth:  o.MinWidth,
		MaxWidth:  o.MaxWidth,
		MinHeight: ClampFloat32(v, o.MinHeight, o.MaxHeight),
		MaxHeight: ClampFloat32(v, o.MinHeight, o.MaxHeight),
	}
}

// A box constraints with the width and height constraints flipped.
func (o BoxConstraints) Flip() BoxConstraints {
	return BoxConstraints{
		MinWidth:  o.MinHeight,
		MaxWidth:  o.MaxHeight,
		MinHeight: o.MinWidth,
		MaxHeight: o.MaxWidth,
	}
}

// Returns box constraints with the same width constraints but with
// unconstrained height.
func (o BoxConstraints) WidthConstraints() BoxConstraints {
	return BoxConstraints{
		MinWidth:  o.MinWidth,
		MaxWidth:  o.MaxWidth,
		MinHeight: 0,
		MaxHeight: SizeInfinity,
	}
}

// Returns box constraints with the same height constraints but with
// unconstrained width.
func (o BoxConstraints) HeightConstraints() BoxConstraints {
	return BoxConstraints{
		MinWidth:  0,
		MaxWidth:  SizeInfinity,
		MinHeight: o.MinHeight,
		MaxHeight: o.MaxHeight,
	}
}

// Returns the width that both satisfies the constraints and is as close as
// possible to the given width.
func (o BoxConstraints) ConstrainWidth(width float32) float32 {
	return ClampFloat32(width, o.MinWidth, o.MaxWidth)
}

// Returns the height that both satisfies the constraints and is as close as
// possible to the given height.
func (o BoxConstraints) ConstrainHeight(height float32) float32 {
	return ClampFloat32(height, o.MinHeight, o.MaxHeight)
}

// Returns the size that both satisfies the constraints and is as close as
// possible to the given size.
func (o BoxConstraints) Constrain(size fyne.Size) fyne.Size {
	return fyne.NewSize(o.ConstrainWidth(size.Width), o.ConstrainHeight(size.Height))
}

// Returns the size that both satisfies the constraints and is as close as
// possible to the given width and height.
func (o BoxConstraints) ConstrainDimensions(width float32, height float32) fyne.Size {
	return fyne.NewSize(o.ConstrainWidth(width), o.ConstrainHeight(height))
}

// The biggest size that satisfies the constraints.
func (o BoxConstraints) MaxSize() fyne.Size {
	return fyne.NewSize(o.ConstrainWidth(SizeInfinity), o.ConstrainHeight(SizeInfinity))
}

// The smallest size that satisfies the constraints.
func (o BoxConstraints) MinSize() fyne.Size {
	return fyne.NewSize(o.ConstrainWidth(0.0), o.ConstrainHeight(0.0))
}

// Whether there is exactly one width value that satisfies the constraints.
func (o BoxConstraints) HasTightWidth() bool {
	return o.MinWidth >= o.MaxWidth
}

/// Whether there is exactly one height value that satisfies the constraints.
func (o BoxConstraints) HasTightHeight() bool {
	return o.MinHeight >= o.MaxHeight
}

func (o BoxConstraints) IsTight() bool {
	return o.HasTightWidth() && o.HasTightHeight()
}

// Returns a size that attempts to meet the following conditions, in order:
func (o BoxConstraints) ConstrainSizeAndAttemptToPreserveAspectRatio(size fyne.Size) fyne.Size {
	if o.IsTight() {
		result := o.MinSize()
		return result
	}

	width := size.Width
	height := size.Height
	aspectRatio := width / height

	if width > o.MaxWidth {
		width = o.MaxWidth
		height = width / aspectRatio
	}

	if height > o.MaxHeight {
		height = o.MaxHeight
		width = height * aspectRatio
	}

	if width < o.MinWidth {
		width = o.MinWidth
		height = width / aspectRatio
	}

	if height < o.MinHeight {
		height = o.MinHeight
		width = height * aspectRatio
	}

	result := o.ConstrainDimensions(width, height)
	return result
}

// Whether there is an upper bound on the maximum width.
func (o BoxConstraints) HasBoundedWidth() bool {
	return o.MaxWidth < SizeInfinity
}

// Whether there is an upper bound on the maximum height.
func (o BoxConstraints) HasBoundedHeight() bool {
	return o.MaxHeight < SizeInfinity
}

// Whether the width constraint is infinite.
//
// Such a constraint is used to indicate that a box should grow as large as
// some other constraint (in this case, horizontally). If constraints are
// infinite, then they must have other (non-infinite) constraints [enforce]d
// upon them, or must be [tighten]ed, before they can be used to derive a
// [Size] for a [RenderBox.size].
func (o BoxConstraints) HasInfiniteWidth() bool {
	return o.MinWidth >= SizeInfinity
}

// Whether the height constraint is infinite.
//
// Such a constraint is used to indicate that a box should grow as large as
// some other constraint (in this case, vertically). If constraints are
// infinite, then they must have other (non-infinite) constraints [enforce]d
// upon them, or must be [tighten]ed, before they can be used to derive a
// [Size] for a [RenderBox.size].
func (o BoxConstraints) HasInfiniteHeight() bool {
	return o.MinHeight >= SizeInfinity
}

// Whether the given size satisfies the constraints.
func (o BoxConstraints) IsSatisfiedBy(size fyne.Size) bool {
	return (o.MinWidth <= size.Width) && (size.Width <= o.MaxWidth) &&
		(o.MinHeight <= size.Height) && (size.Height <= o.MaxHeight)
}

// Scales each constraint parameter by the given factor.
func (o BoxConstraints) Mul(factor float32) BoxConstraints {
	r := o
	if r.MinWidth != SizeInfinity {
		r.MinWidth *= factor
	}
	if r.MaxWidth != SizeInfinity {
		r.MaxWidth *= factor
	}
	if r.MinHeight != SizeInfinity {
		r.MinHeight *= factor
	}
	if r.MaxHeight != SizeInfinity {
		r.MaxHeight *= factor
	}
	return r
}

// Scales each constraint parameter by the inverse of the given factor.
func (o BoxConstraints) Div(factor float32) BoxConstraints {
	r := o
	if r.MinWidth != SizeInfinity {
		r.MinWidth /= factor
	}
	if r.MaxWidth != SizeInfinity {
		r.MaxWidth /= factor
	}
	if r.MinHeight != SizeInfinity {
		r.MinHeight /= factor
	}
	if r.MaxHeight != SizeInfinity {
		r.MaxHeight /= factor
	}
	return r
}

// Returns whether the object's constraints are normalized.
// Constraints are normalized if the minimums are less than or
// equal to the corresponding maximums.
func (o BoxConstraints) IsNormalized() bool {
	return o.MinWidth >= 0.0 &&
		o.MinWidth <= o.MaxWidth &&
		o.MinHeight >= 0.0 &&
		o.MinHeight <= o.MaxHeight
}

// Returns a box constraints that [isNormalized].
//
// The returned [maxWidth] is at least as large as the [minWidth]. Similarly,
// the returned [maxHeight] is at least as large as the [minHeight].
func (o BoxConstraints) Normalize() BoxConstraints {
	r := o
	if o.IsNormalized() {
		return r
	}
	r.MinWidth = 0
	if o.MinWidth >= 0.0 {
		r.MinWidth = o.MinWidth
	}
	r.MinHeight = 0
	if o.MinHeight >= 0.0 {
		r.MinHeight = o.MinHeight
	}
	r.MaxWidth = o.MaxWidth
	if r.MinWidth > o.MaxWidth {
		r.MaxWidth = o.MinWidth
	}
	r.MaxHeight = o.MaxHeight
	if r.MinHeight > o.MaxHeight {
		r.MaxHeight = o.MinHeight
	}
	return r
}

func (o BoxConstraints) Equal(other BoxConstraints) bool {
	return other.MinWidth == o.MinWidth &&
		other.MaxWidth == o.MaxWidth &&
		other.MinHeight == o.MinHeight &&
		other.MaxHeight == o.MaxHeight
}
