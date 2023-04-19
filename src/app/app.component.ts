import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.css']
})
export class AppComponent {
   

    constructor(private _router: Router) {

        // const user = localStorage.getItem('user');
        this._router.navigate(['/auth/login']);
        // if (user) {
        //     this._router.navigate(['/main/home']);
        // } else {
        //     this._router.navigate(['/auth/login']);
        // }

    }
}

