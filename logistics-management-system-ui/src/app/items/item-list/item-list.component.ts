import { Component } from '@angular/core';
import { ItemService } from '../item.service';

const fakeItems = [
    {id: "89819873", name: "bolt", unitPrice: 1.23},
    {id: "82873936", name: "screwdriver", unitPrice: 6.12},
    {id: "28379874", name: "broom handle", unitPrice: 5.34},
    {id: "93849872", name: "mop bucker", unitPrice: 12.99},
    {id: "93849239", name: "cleaning wipes", unitPrice: 4.17},
    {id: "19872364", name: "pipe cleaners", unitPrice: 2.73},
]

@Component({
  selector: 'app-item-list',
  templateUrl: './item-list.component.html',
  styleUrls: ['./item-list.component.scss']
})

export class ItemListComponent {
    public items = fakeItems 

    constructor(private itemService: ItemService){};
}
