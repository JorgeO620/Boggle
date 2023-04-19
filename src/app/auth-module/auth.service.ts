import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, of } from 'rxjs';


@Injectable({
    providedIn: 'root'
})
export class AuthService {
   
    users: any[];

    about = `<p> Hey! my name is Mahamil Gadit. I am currently a senior at the University of Florida 
    looking for jobs. I would like to reach out to specific recruiters at the companies 
    I would like to work for to get more of an insight on the operations of a company.</p>`

    constructor() {

        this.users = [{
            profile_pic: '',
            email: 'mgadit@boggle.com',
            username: 'Mahamil',
            phone: '352-555-5555',
            address: '5225 SW 13th St Gainesville Fl, 32607',
            password: '123456',
            name: 'Mahamil Gadit',
            about: this.about
        }];
    }

    isSameUser(username: string): boolean {
        return this.users.find(u => u.username === username);
    }

    isSameEamil(email: string): boolean {
        return this.users.find(u => u.email === email);
    }

    signup(payload: any): Observable<any> {
        
        if (this.isSameEamil(payload.email)) {
            return of({'error' : 'User with email already exist'});
        }

        if (this.isSameUser(payload.username)) {
            return of({'error' : 'Username must be unique'});
        }

       payload['profile_pic'] = '';
       payload['about'] = this.about;

       this.users.push(payload);

        return of(true);
    }

    login(username: string, password: string): Observable<any> {
        const user = this.users.find(u => u.username === username);

        if (!user || (user && user.password !== password)) {
            return of({'error' : 'Invalid username/password'});
        }

        localStorage.setItem('user', JSON.stringify(user));

        return of(true);
    }
}
