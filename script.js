const preview_image = (event) => {
    const image = document.querySelector('.uploadImage');
    image.src = URL.createObjectURL(event.target.files[0]);
};