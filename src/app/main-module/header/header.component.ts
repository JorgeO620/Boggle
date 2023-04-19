import { Component } from '@angular/core';
import { Router } from '@angular/router';


@Component({
    selector: 'app-header',
    templateUrl: './header.component.html',
    styleUrls: ['./header.component.css']
})
export class HeaderComponent {

    menuItems: any[];
    userProfile: any;
    
    constructor(
        private _router: Router
    ) {
        const user = localStorage.getItem('user');
        this.userProfile = user ? JSON.parse(user) : null;
        
        this.menuItems = [{
            title: 'Home',
            link: '/main/home'
        }, {
            title: 'About Us',
            link: '/main/about-us'
        }, {
            title: 'Connect',
            link: '/main/connect'
        }]
    }

    signOut() {
        localStorage.clear();
        setTimeout(() => {
            this._router.navigate(['/auth/login']);
        }, 1000);
    }
}
