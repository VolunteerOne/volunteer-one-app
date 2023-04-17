// This page is temporarily being used to test other pages (notifications, following, etc.) - matt
import React from 'react';
import { StyleSheet, Dimensions, ScrollView } from 'react-native';
import { Block, theme } from 'galio-framework';
import ProfileItem from '../../components/ProfileItem';
import profile from '../../constants/ProfileTab/profile';

// ================================= View Following Page ================================= //

const { width } = Dimensions.get('screen');

class ViewFollowingPage extends React.Component {
  renderFollowing = () => {
    return (
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.notifications}>
        
        {/* map all notifications using loop */}
        <Block flex>
          {profile.following.followingList.map((profile,i) => (
             <ProfileItem key={i} item={profile} horizontal />
          ))}
        </Block>
    
      </ScrollView>
    )
  }

  render() {
    return (
      <Block flex center style={styles.home}>
        {this.renderFollowing()}
      </Block>
    );
  }
}

const styles = StyleSheet.create({
  home: {
    width: width,    
  },
  notifications: {
    width: width - theme.SIZES.BASE * 2,
    paddingVertical: theme.SIZES.BASE,
  },
});

export default ViewFollowingPage;
