import { Injectable } from '@angular/core';
import { LoginRequest } from './login-request';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Router } from '@angular/router';
import { JwtHelperService } from '@auth0/angular-jwt';
import { TokenResponse } from './token-response';

const helper = new JwtHelperService();

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  constructor(private router: Router, private http: HttpClient) {
    /*const params = new HttpParams({ fromString: window.location.hash.replace('#?', '') });

    const idToken = params.get('token');

    if (idToken) {
      console.log('Setting access token in localStorage...');
      localStorage.setItem('token', idToken);

      this.router.navigate(['/']);
    }*/
  }

  requestToken(request: LoginRequest): Observable<TokenResponse> {
    return this.http.post<TokenResponse>('/auth/login', request);
  }

  login(token: string): any {
    localStorage.setItem('token', token);
  }

  logout() {
    localStorage.removeItem('token');
  }

  isLoggedIn() {
    const token = localStorage.getItem('token');

    return !helper.isTokenExpired(token);
  }

}
