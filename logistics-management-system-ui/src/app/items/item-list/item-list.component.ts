import { Component, OnInit } from '@angular/core';
import { Item, ItemService } from '../item.service';


@Component({
  selector: 'app-item-list',
  templateUrl: './item-list.component.html',
  styleUrls: ['./item-list.component.scss']
})
export class ItemListComponent implements OnInit {
    items: Item[] = []; 

    constructor(private itemService: ItemService){};

    ngOnInit(): void {
        this.itemService.getItems().subscribe(data => {
            this.items = data;
        });
    }

}

