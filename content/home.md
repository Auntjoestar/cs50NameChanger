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

  <label for="name">Tu Nombre:</label>
  <input type="text" id="name" name="name" required class="form-control">
  
  <label for="email">Tu Email:</label>
  <input type="email" id="email" name="email" required class="form-control">

  <label for="review">Tu Reseña:</label><textarea id="review" name="review" rows="5" required class="text-area"></textarea>
  
  <div class="g-recaptcha" data-sitekey="6LcppJQqAAAAABQiUBLKHiszdQD_ajG2qVt5T3F7"></div>

  <button type="submit" class="btn btn-primary">Enviar Reseña</button>
</form>

<script>
  function sendEmail(event) {
    event.preventDefault(); 

    const form = document.getElementById('review-form');
    const formData = {
      name: form.name.value,
      email: form.email.value,
      review: form.review.value,
    };

    emailjs
      .send('service_npss3g4', 'template_ivd82zo', formData)
      .then(() => {
        alert('¡Reseña enviada con éxito! Gracias por tu colaboración.');
        form.reset(); // Limpia el formulario después de enviarlo
      })
      .catch((error) => {
        console.error('Error al enviar el correo:', error);
        alert('Ocurrió un error al enviar tu reseña. Por favor, inténtalo de nuevo.');
      });
  }
</script>