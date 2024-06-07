// Пробная хрень
function sumB() {
      alert("Кролик номер ");  
  }

// Пробная хрень
function sumR() {
    for(let i=1; i<=3; i++) {
      alert("Кролик номер " + i);
    }
  }


// Функция запроса суммы
async function sum(x, y) {
    let result = await fetch(`/sum?x=${x}&y=${y}`)
    let z = await result.text();
    alert("ВЫВОД", x, y);
    return z
}
// Функция запроса разности
async function sub(x, y) {
    let result = await fetch(`/sub?x=${x}&y=${y}`)
    let z = await result.text();
    return z
}
// Функция запроса произведения
async function mult(x, y) {
    let result = await fetch(`/mult?x=${x}&y=${y}`)
    let z = await result.text();
    return z
}
// Функция запроса частного
async function div(x, y) {
    let result = await fetch(`/div?x=${x}&y=${y}`)
    let z = await result.text();
    return z
}