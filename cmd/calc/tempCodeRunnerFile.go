switch data.Operation { // определяем какая операция прошла при нажатии кнопки
	case "sum":
		data.Answer = data.Num1 + data.Num2
		data.IsAnswerExist = true