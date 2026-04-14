let translations = {}

function loadLang(lang) {
  return fetch(`/lang/${lang}.json`)
    .then(res => res.json())
}

function applyLang() {
  document.querySelectorAll("[data-key]").forEach(el => {
    const key = el.dataset.key
    const text = translations[key]

    if (el.tagName === "INPUT") {
      el.placeholder = text
    } else {
      el.innerText = text
    }
  })
}

function setLang(lang) {
  fetch("/api/set-lang", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ lang })
  })

  loadLang(lang).then(data => {
    translations = data
    localStorage.setItem("lang", lang)
    applyLang()
  })
}

function sendData() {
  const name = document.getElementById("name").value

  fetch("/api/post", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ name })
  })
  .then(res => res.json())
  .then(data => {
    document.getElementById("result").innerText = data.reply
  })
}

// โหลดตอนเปิด
const savedLang = localStorage.getItem("lang") || "th"
setLang(savedLang)