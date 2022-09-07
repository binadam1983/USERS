const app = {
    init: () => {
        document.addEventListener('DOMContentLoaded', app.load);
    },
    load: () => {        
        app.getData();
    },
    getData: () => {

        let page = document.body.id;
        switch (page) {
            case 'login-page':
                app.loginPage()
                break;
            case 'welcome-page' || 'user-page':
                app.welcomePage()
                break;
/*             case 'users-page':
                app.welcomePage()
                break; */
            default:
                break;
            }
        },
        
        // This section JS code for / route
        loginPage: () => {
        
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
        },
        //This section contains JS code for the welcome page, first page after successful registration or login of the user
        welcomePage: () => {
            
        let hamburger = document.querySelector('.hamburger');
        let navMenu = document.querySelector('.nav-menu');
        let getUsers = document.querySelector('#get-users');
        const userList = document.querySelector('.users-list');
        const ulList = document.querySelector('.ul-list');
        const frag = document.createDocumentFragment();
        const url = 'http://localhost:3333/user/users';

        const listHeading = document.createElement('h4');
        listHeading.setAttribute('class', 'heading4');
        listHeading.textContent = 'USERS';


        getUsers.addEventListener('click', (e) => {
            e.preventDefault();
            userList.prepend(listHeading);

            fetch(url)
            .then(res => res.json())
            .then(data => {
                for (const email of data) {
                    const li = document.createElement('li');
                    li.textContent = email;
                    frag.appendChild(li);
                }
                console.log(frag)
                ulList.append(frag)
                return;                                    
            })
            .catch(err => console.log(err))
        }); 
        
        hamburger.addEventListener('click', () => {
            hamburger.classList.toggle('active');
            navMenu.classList.toggle('active');
        })
        
        
        document.querySelectorAll('.nav-link').forEach(n => n.addEventListener('click', () => {

           hamburger.classList.remove('active');
           navMenu.classList.remove('active');
        }));        
    }
}
app.init();