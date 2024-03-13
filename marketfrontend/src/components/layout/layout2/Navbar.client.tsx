"use client"

import React, { useState } from "react";
import { Bars3Icon } from "@heroicons/react/24/outline";
import classNames from "classnames";
import Image from "next/image";
import { DropdownMenu, DropdownMenuItem, DropdownMenuTrigger, DropdownMenuContent, DropdownMenuLabel, DropdownMenuSeparator } from "@radix-ui/react-dropdown-menu";

type Props = {
  onMenuButtonClick(): void;
};

type Person = {
  id: number;
  name: string;
};

const people: Person[] = [
  { id: 1, name: 'Nicky Vidovic' },
  { id: 2, name: 'Profile' },
  { id: 3, name: 'Settings' },
  { id: 4, name: 'Extensions' },
  { id: 5, name: 'Dashboard' },
  { id: 6, name: 'Privacy' },
  { id: 7, name: 'Help' },
];

const Navbar = (props: Props) => {
  const [selectedPerson, setSelectedPerson] = useState<Person>(people[0]);

  return (
    <nav
      className={classNames({
        "bg-indigo-700 text-zinc-500": true,
        "flex items-center": true,
        "w-full fixed z-10 px-4 shadow-sm h-16 justify-between": true,
      })}
    >
      <div className="font-bold text-lg bg-white p-2 m-2 rounded-3xl">
       
      </div>
      <div className="font-bold text-lg bg-white p-2 m-2 rounded-3xl">MyTrustElle</div>
      <div className="font-bold text-lg bg-white p-2 m-2 rounded-xl overflow-hidden">
        <DropdownMenu>
          <DropdownMenuTrigger>{selectedPerson?.name}</DropdownMenuTrigger>
          <DropdownMenuContent>
            <DropdownMenuLabel>My Account</DropdownMenuLabel>
            <DropdownMenuSeparator />
            {people.map((person) => (
              <DropdownMenuItem key={person.id} onSelect={() => setSelectedPerson(person)}>
                {person.name}
              </DropdownMenuItem>
            ))}
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
      <button className="lg:hidden md:hidden sm:hidden" onClick={props.onMenuButtonClick}>
        <Bars3Icon className="h-6 w-12 text-white" />
      </button>
    </nav>
  );
};

export default Navbar;
