# LogisticsManagementSystemUi

## Requirements

* Add an Item
  * id number
  * name
  * price

* Remove an Item
  * fails if inventory, orders, or requisitions exist

* Establish allowance
  * item must exist first
  * provide quantity

* Create a stock requisition
  * unique identifier
  * item record must exist
  * item must have allowance

* Initiate stock-fill buy
  * creates stock requisitions for all items under allowance level

* Receive stock requisition
  * sends receipt status to supplier
  * creates pending stow quantity

* Assign location
  * item must exist

* Stow item
  * requires inventory location
  * requires pending stow
  * updates on-hand quantity
  * update pending stow quantity

* Search item
  * show id, name, and price
  * show allowance level, quantity on order, pending stow, and onhand
  * show location

* Search requisition
  * show req id, item ordered, date ordered
  * show quantity ordered, total cost
  * show req status, status date
  * show recieved date or cancellation date

* Event Log
  * append-only transaction ledger
  * every action on item creates event
  * tracks date and any pertinent info

* Initiate inventory
  * places stock items under inventory

* Post inventory
  * updates inventory on-hand
  * creates loss/gain if new quantity differs from current quantity


