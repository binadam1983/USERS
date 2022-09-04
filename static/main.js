// This section JS code for / route

let containerBox = document.querySelector('.container');
let regBox = document.querySelector('.register');
let loginBox = document.querySelector('.login');
let regBackBtn = document.querySelector('.back-reg');
let loginBackBtn = document.querySelector('.back-login');

regBox.addEventListener('click', () => {
    containerBox.classList.add('active-reg');
    
});

regBackBtn.addEventListener('click', (e) => {
    containerBox.classList.remove('active-reg');
    e.stopPropagation();
});


loginBox.addEventListener('click', () => {
    containerBox.classList.add('active-login');

});
loginBackBtn.addEventListener('click', (e) => {
    containerBox.classList.remove('active-login');
    e.stopPropagation();
});

//This section contains JS code for the page after successful registration or login

let hamburger = document.querySelector('.hamburger');
let navMenu = document.querySelector('.nav-menu');

hamburger.addEventListener('click', () => {
    hamburger.classList.toggle('active');
    navMenu.classList.toggle('active');
})


document.querySelectorAll('.nav-link').forEach(n => n.addEventListener('click', () => {
   console.log(document.querySelector('.nav-link'))
   hamburger.classList.remove('active');
   navMenu.classList.remove('active');
}));
