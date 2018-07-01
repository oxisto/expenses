/*
Copyright 2018 Christian Banse

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
