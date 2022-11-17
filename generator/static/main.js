console.log("JS Loaded")

const url = "127.0.0.1:5555"

var inputForm = document.getElementById("inputForm") // получаем форму по айди

inputForm.addEventListener("submit", (e) => { // при нажатии кнопки запускаем передачу данных методом POST

  // предотвращаем автоматическую оправку
  e.preventDefault()

  const formdata = new FormData(inputForm)
  fetch(url, {

    method: "POST",
    body: formdata,
  }).then(
    response => response.text() 
  ).then(
    (data) => { // отправили ответ с сервера
      console.log(data);
      document.getElementById("serverMessageBox").innerHTML = data
    }
  ).catch(
    error => console.error(error)
  )




})
