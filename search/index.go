package search

import (
	"github.com/james-bowman/nlp"
	"github.com/james-bowman/nlp/measures/pairwise"
	"gonum.org/v1/gonum/mat"
)

type Index struct {
	vectoriser  *nlp.CountVectoriser
	transformer *nlp.TfidfTransformer
	reducer     *nlp.TruncatedSVD
	pipeline    *nlp.Pipeline
	lsi         mat.Matrix
}

func NewIndex(removeStopwords bool, k int) *Index {
	vectoriser := nlp.NewCountVectoriser(removeStopwords)
	transformer := nlp.NewTfidfTransformer()
	reducer := nlp.NewTruncatedSVD(k)
	pipeline := nlp.NewPipeline(vectoriser, transformer, reducer)

	return &Index{
		vectoriser:  vectoriser,
		transformer: transformer,
		reducer:     reducer,
		pipeline:    pipeline,
	}
}

func (i *Index) Train(testCorpus []string) error {
	lsi, err := i.pipeline.FitTransform(testCorpus...)
	if err != nil {
		return err
	}
	i.lsi = lsi

	return nil
}

func (index *Index) Query(query string) (int, error) {
	queryVector, err := index.pipeline.Transform(query)
	if err != nil {
		return -1, err
	}

	highestSimilarity := -1.0
	var matched int
	_, docs := index.lsi.Dims()
	for i := 0; i < docs; i++ {
		similarity := pairwise.CosineSimilarity(queryVector.(mat.ColViewer).ColView(0), index.lsi.(mat.ColViewer).ColView(i))
		if similarity > highestSimilarity {
			matched = i
			highestSimilarity = similarity
		}
	}

	return matched, nil
}
