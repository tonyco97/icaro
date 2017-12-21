/*
 * Copyright (C) 2017 Nethesis S.r.l.
 * http://www.nethesis.it - info@nethesis.it
 *
 * This file is part of Icaro project.
 *
 * NethServer is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License,
 * or any later version.
 *
 * NethServer is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with NethServer.  If not, see COPYING.
 *
 * author: Edoardo Spadoni <edoardo.spadoni@nethesis.it>
 */

package models

import "time"

type Unit struct {
	Id          int       `db:"id" json:"id"`
	HotspotId   int       `db:"hotspot_id" json:"hotspot_id"`
	MacAddress  string    `db:"mac_address" json:"mac_address"`
	Description string    `db:"description" json:"description"`
	Uuid        string    `db:"uuid" json:"uuid"`
	Created     time.Time `db:"created" json:"created"`
}