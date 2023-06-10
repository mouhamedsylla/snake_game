package main

type SneakBody struct {
	Parts  []SnakePart
	Xspeed int
	Yspeed int
}

type SnakePart struct {
	X int
	Y int
}

func (sb *SneakBody) ChangeDir(vertical, horizontal int) {

	sb.Xspeed = horizontal
	sb.Yspeed = vertical
}

func (sb *SneakBody) Update(width, height int) {

	sb.Parts = append(sb.Parts, sb.Parts[len(sb.Parts)-1].GetUpdatePart(sb, width, height))
	sb.Parts = sb.Parts[1:]
}

func (sp *SnakePart) GetUpdatePart(sb *SneakBody, width, height int) SnakePart {

	newPart := *sp

	newPart.X = (newPart.X + sb.Xspeed) % width
	if newPart.X < 0 {
		newPart.X += width
	}

	newPart.Y = (newPart.Y + sb.Yspeed) % height
	if newPart.Y < 0 {
		newPart.Y += height
	}

	return newPart
}
