+++
draft = false
title = "Bienvenido"
+++

## Bienvenido

Esta es la página de la aplicación CS50 Name Changer, haz click [aquí](./downloads) para acceder al link de descarga. 

---

## Envía tu reseña
¡Nos encantaría conocer tu opinión sobre la aplicación! Por favor, completa el formulario a continuación para compartir tu experiencia:
<form id="review-form" onsubmit="sendEmail(event)">
  <p id="result" hidden></p>
  <label for="name">Tu Nombre:</label>
  <input type="text" id="name" name="name" required class="form-control">
  
  <label for="email">Tu Email:</label>
  <input type="email" id="email" name="email" required class="form-control">

  <label for="review">Tu Reseña:</label>
  <textarea id="review" name="review" rows="5" required class="text-area"></textarea>
  
  <!-- reCAPTCHA -->
  <div class="g-recaptcha" data-sitekey="6LcppJQqAAAAABQiUBLKHiszdQD_ajG2qVt5T3F7"></div>

  <button type="submit" class="btn btn-primary">Enviar Reseña</button>
</form>

<script>
  function sendEmail(e) {
    e.preventDefault(); 
    
    const result = document.getElementById("result");

    const recaptchaResponse = grecaptcha.getResponse();
    if (!recaptchaResponse) {
      result.innerText = "Por favor, completa el CAPTCHA antes de enviar el formulario.";
      result.style.color = "#CC0000";
      result.hidden = false;
      result.scrollIntoView({ behavior: 'smooth', block: 'start' });
      return;
    }

    const form = document.getElementById('review-form');

    const formData = {
      name: form.name.value,
      email: form.email.value,
      review: form.review.value,
      recaptcha: recaptchaResponse, 
    };

    emailjs
      .send('service_npss3g4', 'template_ivd82zo', formData)
      .then(() => {
        result.innerText = "Review enviada exitosamente";
        result.style.color = "#4BB543";
        result.hidden = false;
        result.scrollIntoView({ behavior: 'smooth', block: 'start' });
        grecaptcha.reset(); 
        form.reset(); 
      })
      .catch((error) => {
        result.innerText = `Error: ${error.toString()}`
        result.style.color = "#CC0000";
        result.hidden = false;
        result.scrollIntoView({ behavior: 'smooth', block: 'start' });
      });
  }
</script>
