// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

contract EcoTrace {
    address public owner;
    mapping(string => EcoInfo) public tracedata;

    struct EcoInfo {
        string ecode;
        string codata;
        string operator;
        string waterdata;
    }

    constructor() {
        owner = msg.sender;
    }

    function addOrupdateProdData(
        string memory id,
        string memory ecode,
        string memory codata,
        string memory operator,
        string memory waterdata
    ) external {
        tracedata[id].ecode = ecode;
        tracedata[id].codata = codata;
        tracedata[id].operator = operator;
        tracedata[id].waterdata = waterdata;
    }

    function tracedataById(string memory id)
        external
        view
        returns (EcoInfo memory)
    {
        return tracedata[id];
    }
}
