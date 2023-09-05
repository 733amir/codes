package main

func main() {
}

type ListNode struct {
    Val int
    Next *ListNode
}

// func middleNode(head *ListNode) *ListNode {
//     h1, h2, biJump := head, head, true
//
//     for h1.Next != nil {
//         h1 = h1.Next
//
//         biJump = !biJump
//         if biJump {
//             h2 = h2.Next
//         }
//     }
//
//     if !biJump {
//         h2 = h2.Next
//     }
//
//     return h2
// }

func middleNode(head *ListNode) *ListNode {
    h1, h2 := head, head

    for h1.Next != nil {
        h1 = h1.Next
        if h1.Next != nil {
            h1 = h1.Next
        }
        h2 = h2.Next
    }

    return h2
}
