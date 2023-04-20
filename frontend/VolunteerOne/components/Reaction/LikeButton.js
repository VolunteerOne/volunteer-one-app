
import React from "react";
import { Block, Button, Text } from "galio-framework";
import { StyleSheet, TouchableOpacity} from "react-native";
import Icon from "../Icon";

class LikeButton extends React.Component {
    state = {
        count: this.props.likeCount
    }

    render() {    
        const incrementLikes = () => {
            console.log("Like button pressed")
            let newCount = this.state.count + 1
            this.setState ({
                count: newCount
            })
        }

        return (
            <Block style={styles.block}>
            <TouchableOpacity style={styles.button}
            onPress={incrementLikes} >      
      <Block flex row>
                <Icon
      family="MaterialIcons"
      size={15}
      name="thumb-up"
      color="#32325D"
    />
            {this.state.count == 0 ? 
            <Text style={styles.titleText}>Like</Text>
            :
            <Text style={styles.titleText}>{this.state.count} Likes</Text>  
            }

            </Block>
          </TouchableOpacity>
          </Block>
        );
        }
        
}
const styles = StyleSheet.create({
    block: {
    //    height: '100%'
    },
    titleText: {
        fontSize: 15,
        color: "#32325D",
        marginLeft: 5,
        marginTop: -1
      },
    button: {
        borderless: true,
        backgroundColor: '#fff',
        shadowColor: '#fff',
        height: 20,
        marginTop: -30,
        marginLeft: 27,
        zIndex : 1,
        
    }

});

export default LikeButton;
    