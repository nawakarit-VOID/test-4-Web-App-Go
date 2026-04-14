let currentLang = "th"
let translations = {}

function loadLang(lang) {
  return fetch(`/lang/${lang}.json`)
    .then(res => res.json())
}

function applyLang() {
  console.log("translations:", translations)

  document.querySelectorAll("[data-key]").forEach(el => {
    const key = el.dataset.key
    const text = translations[key]

    console.log("key:", key)
    console.log("text:", text)

    if (el.tagName === "INPUT") {
      el.placeholder = text
    } else {
      el.innerText = text
    }
  })
}

function setLang(lang) {
  currentLang = lang

  loadLang(lang).then(data => {
    translations = data

    // เก็บไว้
    localStorage.setItem("lang", lang)

    applyLang()
  })
}

// โหลดครั้งแรก
const savedLang = localStorage.getItem("lang") || "th"
setLang(savedLang)