import React, { Fragment, useState } from "react";
import { Bars3Icon } from "@heroicons/react/24/outline";
import classNames from "classnames";
import Image from "next/image";
import KijijiLogo from "/public/KijijiLogo.png";
import { Listbox } from '@headlessui/react';
import { DropdownMenu, DropdownMenuItem, DropdownMenuTrigger, DropdownMenuContent, DropdownMenuLabel, DropdownMenuSeparator } from "@radix-ui/react-dropdown-menu";

type Props = {
  onMenuButtonClick(): void;
};

type Person = {
  id: number;
  name: string;
};

const people: Person[] = [
  { id: 1, name: 'Lapointe Group' },
  { id: 2, name: 'Nick Group' },
  { id: 3, name: 'George Group' },
  { id: 4, name: 'Sheila Group' },
  { id: 5, name: 'Hamburger Group' },
  { id: 6, name: '2Lapointe Group' },
  { id: 7, name: '3Nick Group' },
  { id: 8, name: '4George Group' },
  { id: 9, name: '5Sheila Group' },
  { id: 10, name: '6Hamburger Group' },
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
        <Image
          src={KijijiLogo}
          height={66}
          width={66}
          alt="profile image"
          className="rounded-full"
        />
      </div>
      <div className="font-bold text-lg bg-white p-2 m-2 rounded-3xl">KCM</div>
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
