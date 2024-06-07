// Функция запроса суммы
async function sum(x, y) {
  var x = Number(document.getElementById("num1").value);  // получаем числа х, у с формы
  var y = Number(document.getElementById("num2").value);
   let result = await fetch(`/sum?x=${x}&y=${y}`)         // передаем параметры через querysrting и ждем ответа от /sum
   let z = await result.text();                           // получает ответ - объект (текст)
   document.getElementById("result").innerHTML = z;       // выводим в форму 
   //return z                                             // TODO: пока не понял зачем возвращать z - работает и так
}
// Функция запроса разности
async function sub(x, y) {
  var x = Number(document.getElementById("num1").value);   
  var y = Number(document.getElementById("num2").value);
  let result = await fetch(`/sub?x=${x}&y=${y}`)
  let z = await result.text();
  document.getElementById("result").innerHTML = z;
  return z
}
// Функция запроса произведения
async function mult(x, y) {
  var x = Number(document.getElementById("num1").value);   
  var y = Number(document.getElementById("num2").value);
  let result = await fetch(`/mult?x=${x}&y=${y}`)
  let z = await result.text();
  document.getElementById("result").innerHTML = z;
  return z
}
// Функция запроса частного
async function div(x, y) {
  var x = Number(document.getElementById("num1").value);   
  var y = Number(document.getElementById("num2").value);
  let result = await fetch(`/div?x=${x}&y=${y}`)
  let z = await result.text();
  document.getElementById("result").innerHTML = z;
  return z
}
