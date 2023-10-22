import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface Item {
    id: string
    name: string
    unitPrice: number
}

@Injectable({
  providedIn: 'root'
})
export class ItemService {

    constructor(private http: HttpClient) {}

  getItems(): Observable<Item[]> {
    return this.http.get<Item[]>("http://localhost:3000/api/items");
  }
}


