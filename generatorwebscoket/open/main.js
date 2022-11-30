(function () {
var webSocket = new WebSocket('ws://127.0.0.1:5555/generator');
webSocket.addEventListener('message', (event) => {  // обработчик response
  var response = JSON.parse(event.data);
  if (!response) return;
  document.getElementById('serverMessageBox').innerHTML = response.result.join(" "); // вывод результата генерации
});
webSocket.addEventListener('close', (event) => {  // обработчик закрытия соденинения 
  document.querySelectorAll('.send').forEach((element) => {
    element.disabled = false;
  });
});
	if (typeof document !== "undefined") {  // если документ существует, ждём его полной загрузки
    document.addEventListener("DOMContentLoaded", function (event) {
  var formElement = document.getElementById("inputForm"); // находим форму
  if (!formElement) {
    return console.error("Form element not found");
  }
  formElement.addEventListener("submit", function (event) {// при отправке формы парсим данные с неё и передаём на сервер
    event.preventDefault();
    var data = new FormData(event.target);
    var flows = parseInt(data.get("flows"));
    var maxNumber = parseInt(data.get("maxNumber"));
    webSocket.send(JSON.stringify({
      "flows": flows,
      "maxNumber": maxNumber,
    }));
  });
});
};
})();


