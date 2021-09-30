package internal

type SerializeRecord  struct{
	Path         	[]string
	IsFoLder        []bool
	Content 		[][]byte
	Size			int
}

func (sr *SerializeRecord) SetSerializeRecord(node *Component)  *SerializeRecord{
	sr.Path = append(sr.Path, node.Path)
	sr.IsFoLder = append(sr.IsFoLder, node.IsFolder)
	sr.Content = append(sr.Content, node.Content)
	sr.Size +=1
	return sr
}

func (sr *SerializeRecord) SerializeSonMatch(node *Component) {
	for _, v := range node.SonComponent {
		if v != nil {
			sr.SetSerializeRecord(v)
			sr.SerializeSonMatch(v)
		}
	}

}
