package domain

type Calculation interface {
  Calculate() (int64, error)
}

type Const int64

func (cnst Const) Calculate() (int64, error) {
  return int64(cnst), nil
}
