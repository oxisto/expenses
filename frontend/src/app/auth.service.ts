import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { JwtHelperService } from '@auth0/angular-jwt';
import { EMPTY, Observable, ObservableInput } from 'rxjs';
import { LoginRequest } from './login-request';
import { TokenResponse } from './token-response';
import { User } from './user';

export function getToken() {
  return localStorage.getItem('token');
}

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

  handleHttpError(err: HttpErrorResponse): ObservableInput<any> {
    if (err.status === 401) {
      this.logout();
    }

    return EMPTY;
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

  getUser(): User {
    const token = getToken();

    return helper.decodeToken(token).user;
  }

  isLoggedIn() {
    const token = getToken();

    return !helper.isTokenExpired(token);
  }

}
